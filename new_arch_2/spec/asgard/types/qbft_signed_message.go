package types

const (
	ProposalMessageType = iota
	PrepareMessageType
	CommitMessageType
	RoundChangeMessageType
)

type QBFTMessage struct {
	MsgType uint64
	Round   uint64 // QBFT round for which the msg is for

	Root                     [32]byte `ssz-size:"32"`
	DataRound                uint64
	RoundChangeJustification [][]byte `ssz-max:"13,65536"` // 2^16
	PrepareJustification     [][]byte `ssz-max:"13,65536"` // 2^16
}

type QBFTSignedMessage struct {
	// Message is at the top for quick identifier look (see docs)
	Message   QBFTMessage
	Signature [96]byte `ssz-size:"96"`
	Signers   []uint64 `ssz-max:"13"`
	FullData  []byte   `ssz-max:"4259840"`
}
