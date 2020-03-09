// +build devnet

package constants

const (
	PostCashOutDelayBlock 	= 60*5
	VoteCashOutDelayBlock = PostCashOutDelayBlock
	VoteRegenerateTime 		= 60*5
	PowerDownBlockInterval 	= 60*5
	MinEpochDuration 		= 100

	StakeFreezeTime      	= 60*5
	WindowSize           = 60 * 60 * 24

	PerTicketPrice = 1
	PerTicketPriceStr = "1.000000"
	PerTicketWeight = uint64(1e7)

	ClientName              = "Cos-go-devnet"

	// vest delegation
	MinVestDelegationInBlocks = 5 * 60 / BlockInterval			// 5 minutes
	VestDelegationDeliveryInBlocks = 5 * 60 / BlockInterval		// 5 minutes
)

// hard forks
const (
	Original uint64 = 0
	HardFork1 uint64 = 3600
	HardFork2 uint64 = 8330951
	HardFork3 uint64 = 99999999	// TODO: SET CORRECT BLOCK NUMBER
)
