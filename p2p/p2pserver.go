package p2p

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/coschain/contentos-go/p2p/common"
	"github.com/coschain/contentos-go/p2p/message/msg_pack"
	comm "github.com/coschain/contentos-go/p2p/depend/common"
	"github.com/coschain/contentos-go/p2p/depend/common/config"
	"github.com/coschain/contentos-go/p2p/depend/common/log"
	msgtypes "github.com/coschain/contentos-go/p2p/message/types"
	"github.com/coschain/contentos-go/p2p/message/utils"
	"github.com/coschain/contentos-go/p2p/net/netserver"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/p2p/peer"
	"github.com/coschain/contentos-go/node"
)

//P2PServer control all network activities
type P2PServer struct {
	Network   iservices.P2P
	msgRouter *utils.MessageRouter
	ReconnectAddrs
	recentPeers    map[uint32][]string
	quitSyncRecent chan bool
	quitOnline     chan bool
	quitHeartBeat  chan bool
}

//ReconnectAddrs contain addr need to reconnect
type ReconnectAddrs struct {
	sync.RWMutex
	RetryAddrs map[string]int
}

//NewServer return a new p2pserver according to the pubkey
func NewServer(ctx *node.ServiceContext) (*P2PServer, error) {
	n := netserver.NewNetServer(ctx)

	p := &P2PServer{
		Network: n,
	}

	p.msgRouter = utils.NewMsgRouter(p.Network)
	p.recentPeers = make(map[uint32][]string)
	p.quitSyncRecent = make(chan bool)
	p.quitOnline = make(chan bool)
	p.quitHeartBeat = make(chan bool)
	return p, nil
}

//GetConnectionCnt return the established connect count
func (this *P2PServer) GetConnectionCnt() uint32 {
	return this.Network.GetConnectionCnt()
}

//Start create all services
func (this *P2PServer) Start(node *node.Node) error {
	if this.Network != nil {
		this.Network.Start(node)
	} else {
		return errors.New("[p2p]network invalid")
	}
	if this.msgRouter != nil {
		this.msgRouter.Start()
	} else {
		return errors.New("[p2p]msg router invalid")
	}
	this.tryRecentPeers()
	go this.connectSeedService()
	go this.syncUpRecentPeers()
	go this.keepOnlineService()
	go this.heartBeatService()
	return nil
}

//Stop halt all service by send signal to channels
func (this *P2PServer) Stop() error {
	this.Network.Halt()
	this.quitSyncRecent <- true
	this.quitOnline <- true
	this.quitHeartBeat <- true
	this.msgRouter.Stop()
	return nil
}

// GetNetWork returns the low level netserver
func (this *P2PServer) GetNetWork() iservices.P2P {
	return this.Network
}

//GetPort return two network port
func (this *P2PServer) GetPort() (uint16, uint16) {
	return this.Network.GetSyncPort(), this.Network.GetConsPort()
}

//GetVersion return self version
func (this *P2PServer) GetVersion() uint32 {
	return this.Network.GetVersion()
}

//GetNeighborAddrs return all nbr`s address
func (this *P2PServer) GetNeighborAddrs() []common.PeerAddr {
	return this.Network.GetNeighborAddrs()
}

//Send tranfer buffer to peer
func (this *P2PServer) Send(p *peer.Peer, msg msgtypes.Message,
	isConsensus bool) error {
	if this.Network.IsPeerEstablished(p) {
		return this.Network.Send(p, msg, isConsensus)
	}
	log.Warnf("[p2p]send to a not ESTABLISH peer %d",
		p.GetID())
	return errors.New("[p2p]send to a not ESTABLISH peer")
}

// GetID returns local node id
func (this *P2PServer) GetID() uint64 {
	return this.Network.GetID()
}


// Todo: remove it if no use
func (this *P2PServer) GetConnectionState() uint32 {
	return common.INIT
}

//GetTime return lastet contact time
func (this *P2PServer) GetTime() int64 {
	return this.Network.GetTime()
}


//blockSyncFinished compare all nbr peers and self height at beginning
func (this *P2PServer) blockSyncFinished() bool {
	peers := this.Network.GetNeighbors()
	if len(peers) == 0 {
		return false
	}

	//blockHeight := this.ledger.GetCurrentBlockHeight()

	blockHeight := uint32(0)

	for _, v := range peers {
		if blockHeight < uint32(v.GetHeight()) {
			return false
		}
	}
	return true
}

//WaitForSyncBlkFinish compare the height of self and remote peer in loop
func (this *P2PServer) WaitForSyncBlkFinish() {
	consensusType := strings.ToLower(config.DefConfig.Genesis.ConsensusType)
	if consensusType == "solo" {
		return
	}

	for {
		//headerHeight := this.ledger.GetCurrentHeaderHeight()
		//currentBlkHeight := this.ledger.GetCurrentBlockHeight()
		//log.Info("[p2p]WaitForSyncBlkFinish... current block height is ",
		//	currentBlkHeight, " ,current header height is ", headerHeight)

		if this.blockSyncFinished() {
			break
		}

		<-time.After(time.Second * (time.Duration(common.SYNC_BLK_WAIT)))
	}
}

//WaitForPeersStart check whether enough peer linked in loop
func (this *P2PServer) WaitForPeersStart() {
	periodTime := config.DEFAULT_GEN_BLOCK_TIME / common.UPDATE_RATE_PER_BLOCK
	for {
		log.Info("[p2p]Wait for minimum connection...")
		if this.reachMinConnection() {
			break
		}

		<-time.After(time.Second * (time.Duration(periodTime)))
	}
}

//connectSeeds connect the seeds in seedlist and call for nbr list
func (this *P2PServer) connectSeeds() {
	seedNodes := make([]string, 0)
	pList := make([]*peer.Peer, 0)
	for _, n := range config.DefConfig.Genesis.SeedList {
		ip, err := common.ParseIPAddr(n)
		if err != nil {
			log.Warnf("[p2p]seed peer %s address format is wrong", n)
			continue
		}
		ns, err := net.LookupHost(ip)
		if err != nil {
			log.Warnf("[p2p]resolve err: %s", err.Error())
			continue
		}
		port, err := common.ParseIPPort(n)
		if err != nil {
			log.Warnf("[p2p]seed peer %s address format is wrong", n)
			continue
		}
		seedNodes = append(seedNodes, ns[0]+port)
	}

	for _, nodeAddr := range seedNodes {
		var ip net.IP
		np := this.Network.GetNp()
		np.Lock()
		for _, tn := range np.List {
			ipAddr, _ := tn.GetAddr16()
			ip = ipAddr[:]
			addrString := ip.To16().String() + ":" +
				strconv.Itoa(int(tn.GetSyncPort()))
			if nodeAddr == addrString && tn.GetSyncState() == common.ESTABLISH {
				pList = append(pList, tn)
			}
		}
		np.Unlock()
	}
	if len(pList) > 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(pList))
		this.reqNbrList(pList[index])
	} else { //not found
		for _, nodeAddr := range seedNodes {
			go this.Network.Connect(nodeAddr, false)
		}
	}
}

//reachMinConnection return whether net layer have enough link under different config
func (this *P2PServer) reachMinConnection() bool {
	if config.DefConfig.Consensus.EnableConsensus == false {
		//just sync
		return true
	}
	consensusType := strings.ToLower(config.DefConfig.Genesis.ConsensusType)
	if consensusType == "" {
		consensusType = "dbft"
	}
	minCount := config.DBFT_MIN_NODE_NUM
	switch consensusType {
	case "dbft":
	case "solo":
		minCount = config.SOLO_MIN_NODE_NUM
	case "vbft":
		minCount = config.VBFT_MIN_NODE_NUM

	}
	return int(this.GetConnectionCnt())+1 >= minCount
}

//getNode returns the peer with the id
func (this *P2PServer) getNode(id uint64) *peer.Peer {
	return this.Network.GetPeer(id)
}

//retryInactivePeer try to connect peer in INACTIVITY state
func (this *P2PServer) retryInactivePeer() {
	np := this.Network.GetNp()
	np.Lock()
	var ip net.IP
	neighborPeers := make(map[uint64]*peer.Peer)
	for _, p := range np.List {
		addr, _ := p.GetAddr16()
		ip = addr[:]
		nodeAddr := ip.To16().String() + ":" +
			strconv.Itoa(int(p.GetSyncPort()))
		if p.GetSyncState() == common.INACTIVITY {
			log.Debugf("[p2p] try reconnect %s", nodeAddr)
			//add addr to retry list
			this.addToRetryList(nodeAddr)
			p.CloseSync()
			p.CloseCons()
		} else {
			//add others to tmp node map
			this.removeFromRetryList(nodeAddr)
			neighborPeers[p.GetID()] = p
		}
	}

	np.List = neighborPeers
	np.Unlock()

	connCount := uint(this.Network.GetOutConnRecordLen())
	if connCount >= config.DefConfig.P2PNode.MaxConnOutBound {
		log.Warnf("[p2p]Connect: out connections(%d) reach the max limit(%d)", connCount,
			config.DefConfig.P2PNode.MaxConnOutBound)
		return
	}

	//try connect
	if len(this.RetryAddrs) > 0 {
		this.ReconnectAddrs.Lock()

		list := make(map[string]int)
		addrs := make([]string, 0, len(this.RetryAddrs))
		for addr, v := range this.RetryAddrs {
			v += 1
			addrs = append(addrs, addr)
			if v < common.MAX_RETRY_COUNT {
				list[addr] = v
			}
			if v >= common.MAX_RETRY_COUNT {
				this.Network.RemoveFromConnectingList(addr)
				remotePeer := this.Network.GetPeerFromAddr(addr)
				if remotePeer != nil {
					if remotePeer.SyncLink.GetAddr() == addr {
						this.Network.RemovePeerSyncAddress(addr)
						this.Network.RemovePeerConsAddress(addr)
					}
					if remotePeer.ConsLink.GetAddr() == addr {
						this.Network.RemovePeerConsAddress(addr)
					}
					this.Network.DelNbrNode(remotePeer.GetID())
				}
			}
		}

		this.RetryAddrs = list
		this.ReconnectAddrs.Unlock()
		for _, addr := range addrs {
			rand.Seed(time.Now().UnixNano())
			log.Debug("[p2p]Try to reconnect peer, peer addr is ", addr)
			<-time.After(time.Duration(rand.Intn(common.CONN_MAX_BACK)) * time.Millisecond)
			log.Debug("[p2p]Back off time`s up, start connect node")
			this.Network.Connect(addr, false)
		}

	}
}

//connectSeedService make sure seed peer be connected
func (this *P2PServer) connectSeedService() {
	t := time.NewTimer(time.Second * common.CONN_MONITOR)
	for {
		select {
		case <-t.C:
			this.connectSeeds()
			t.Stop()
			if this.reachMinConnection() {
				t.Reset(time.Second * time.Duration(10*common.CONN_MONITOR))
			} else {
				t.Reset(time.Second * common.CONN_MONITOR)
			}
		case <-this.quitOnline:
			t.Stop()
			break
		}
	}
}

//keepOnline try connect lost peer
func (this *P2PServer) keepOnlineService() {
	t := time.NewTimer(time.Second * common.CONN_MONITOR)
	for {
		select {
		case <-t.C:
			this.retryInactivePeer()
			t.Stop()
			t.Reset(time.Second * common.CONN_MONITOR)
		case <-this.quitOnline:
			t.Stop()
			break
		}
	}
}

//reqNbrList ask the peer for its neighbor list
func (this *P2PServer) reqNbrList(p *peer.Peer) {
	msg := msgpack.NewAddrReq()
	go this.Send(p, msg, false)
}

//heartBeat send ping to nbr peers and check the timeout
func (this *P2PServer) heartBeatService() {
	var periodTime uint
	periodTime = config.DEFAULT_GEN_BLOCK_TIME / common.UPDATE_RATE_PER_BLOCK
	t := time.NewTicker(time.Second * (time.Duration(periodTime)))

	for {
		select {
		case <-t.C:
			this.ping()
			this.timeout()
		case <-this.quitHeartBeat:
			t.Stop()
			break
		}
	}
}

//ping send pkg to get pong msg from others
func (this *P2PServer) ping() {
	peers := this.Network.GetNeighbors()
	this.pingTo(peers)
}

//pings send pkgs to get pong msg from others
func (this *P2PServer) pingTo(peers []*peer.Peer) {
	for _, p := range peers {
		if p.GetSyncState() == common.ESTABLISH {

			//height := this.ledger.GetCurrentBlockHeight()

			height := 0

			ping := msgpack.NewPingMsg(uint64(height))
			go this.Send(p, ping, false)
		}
	}
}

//timeout trace whether some peer be long time no response
func (this *P2PServer) timeout() {
	peers := this.Network.GetNeighbors()
	var periodTime uint
	periodTime = config.DEFAULT_GEN_BLOCK_TIME / common.UPDATE_RATE_PER_BLOCK
	for _, p := range peers {
		if p.GetSyncState() == common.ESTABLISH {
			t := p.GetContactTime()
			if t.Before(time.Now().Add(-1 * time.Second *
				time.Duration(periodTime) * common.KEEPALIVE_TIMEOUT)) {
				log.Warnf("[p2p]keep alive timeout!!!lost remote peer %d - %s from %s", p.GetID(), p.SyncLink.GetAddr(), t.String())
				p.CloseSync()
				p.CloseCons()
			}
		}
	}
}

//addToRetryList add retry address to ReconnectAddrs
func (this *P2PServer) addToRetryList(addr string) {
	this.ReconnectAddrs.Lock()
	defer this.ReconnectAddrs.Unlock()
	if this.RetryAddrs == nil {
		this.RetryAddrs = make(map[string]int)
	}
	if _, ok := this.RetryAddrs[addr]; ok {
		delete(this.RetryAddrs, addr)
	}
	//alway set retry to 0
	this.RetryAddrs[addr] = 0
}

//removeFromRetryList remove connected address from ReconnectAddrs
func (this *P2PServer) removeFromRetryList(addr string) {
	this.ReconnectAddrs.Lock()
	defer this.ReconnectAddrs.Unlock()
	if len(this.RetryAddrs) > 0 {
		if _, ok := this.RetryAddrs[addr]; ok {
			delete(this.RetryAddrs, addr)
		}
	}
}

//tryRecentPeers try connect recent contact peer when service start
func (this *P2PServer) tryRecentPeers() {
	netID := config.DefConfig.P2PNode.NetworkMagic
	if comm.FileExisted(common.RECENT_FILE_NAME) {
		buf, err := ioutil.ReadFile(common.RECENT_FILE_NAME)
		if err != nil {
			log.Warn("[p2p]read %s fail:%s, connect recent peers cancel", common.RECENT_FILE_NAME, err.Error())
			return
		}

		err = json.Unmarshal(buf, &this.recentPeers)
		if err != nil {
			log.Warn("[p2p]parse recent peer file fail: ", err)
			return
		}
		if len(this.recentPeers[netID]) > 0 {
			log.Info("[p2p]try to connect recent peer")
		}
		for _, v := range this.recentPeers[netID] {
			go this.Network.Connect(v, false)
		}

	}
}

//syncUpRecentPeers sync up recent peers periodically
func (this *P2PServer) syncUpRecentPeers() {
	periodTime := common.RECENT_TIMEOUT
	t := time.NewTicker(time.Second * (time.Duration(periodTime)))
	for {
		select {
		case <-t.C:
			this.syncPeerAddr()
		case <-this.quitSyncRecent:
			t.Stop()
			break
		}
	}

}

//syncPeerAddr compare snapshot of recent peer with current link,then persist the list
func (this *P2PServer) syncPeerAddr() {
	changed := false
	netID := config.DefConfig.P2PNode.NetworkMagic
	for i := 0; i < len(this.recentPeers[netID]); i++ {
		p := this.Network.GetPeerFromAddr(this.recentPeers[netID][i])
		if p == nil || (p != nil && p.GetSyncState() != common.ESTABLISH) {
			this.recentPeers[netID] = append(this.recentPeers[netID][:i], this.recentPeers[netID][i+1:]...)
			changed = true
			i--
		}
	}
	left := common.RECENT_LIMIT - len(this.recentPeers[netID])
	if left > 0 {
		np := this.Network.GetNp()
		np.Lock()
		var ip net.IP
		for _, p := range np.List {
			addr, _ := p.GetAddr16()
			ip = addr[:]
			nodeAddr := ip.To16().String() + ":" +
				strconv.Itoa(int(p.GetSyncPort()))
			found := false
			for i := 0; i < len(this.recentPeers[netID]); i++ {
				if nodeAddr == this.recentPeers[netID][i] {
					found = true
					break
				}
			}
			if !found {
				this.recentPeers[netID] = append(this.recentPeers[netID], nodeAddr)
				left--
				changed = true
				if left == 0 {
					break
				}
			}
		}
		np.Unlock()
	} else {
		if left < 0 {
			left = -left
			this.recentPeers[netID] = append(this.recentPeers[netID][:0], this.recentPeers[netID][0+left:]...)
			changed = true
		}

	}
	if changed {
		buf, err := json.Marshal(this.recentPeers)
		if err != nil {
			log.Warn("[p2p]package recent peer fail: ", err)
			return
		}
		err = ioutil.WriteFile(common.RECENT_FILE_NAME, buf, os.ModePerm)
		if err != nil {
			log.Warn("[p2p]write recent peer fail: ", err)
		}
	}
}
