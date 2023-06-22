package qbft

import "ssv-experiments/new_arch/p2p"

const (
	ProposalMessageType = iota
	PrepareMessageType
	CommitMessageType
	RoundChangeMessageType
)

type Message struct {
	MsgType    uint64
	Round      uint64         // QBFT round for which the msg is for
	Identifier p2p.Identifier `ssz-size:"56"` // instance Identifier this msg belongs to

	Root                     [32]byte `ssz-size:"32"`
	DataRound                uint64
	RoundChangeJustification [][]byte `ssz-max:"13,65536"` // 2^16
	PrepareJustification     [][]byte `ssz-max:"13,65536"` // 2^16
}

func (msg *Message) GetIdentifier() [56]byte {
	return msg.Identifier
}

type SignedMessage struct {
	Signature [96]byte `ssz-size:"96"`
	Signers   []uint64 `ssz-max:"13"`
	Message   Message
	FullData  []byte `ssz-max:"4259840"`
}

func (signed *SignedMessage) GetIdentifier() [56]byte {
	return signed.Message.Identifier
}
