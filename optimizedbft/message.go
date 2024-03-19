package optimizedbft

import "github.com/herumi/bls-eth-go-binary/bls"

type MessageType = uint64

const (
	Propose MessageType = iota
	Prepare
	Commit
)

type Message struct {
	Type       MessageType
	OperatorID uint64
	Height     uint64
	Round      uint64
	Data       []byte

	PartialSignature *bls.Sign
}
