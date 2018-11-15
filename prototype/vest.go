package prototype

import "github.com/coschain/contentos-go/common/encoding"

func (m *Vest) OpeEncode() ([]byte, error) {
	return encoding.Encode(m.Value)
}

func MakeVest(value uint64) *Vest {
	return &Vest{Value:value}
}