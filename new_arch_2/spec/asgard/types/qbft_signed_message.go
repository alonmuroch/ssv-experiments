package types

import "github.com/pkg/errors"

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

func (msg *QBFTMessage) GetRoundChangeJustifications() ([]*QBFTSignedMessage, error) {
	return unmarshalJustifications(msg.RoundChangeJustification)
}

func (msg *QBFTMessage) GetPrepareJustifications() ([]*QBFTSignedMessage, error) {
	return unmarshalJustifications(msg.PrepareJustification)
}

func unmarshalJustifications(data [][]byte) ([]*QBFTSignedMessage, error) {
	ret := make([]*QBFTSignedMessage, len(data))
	for i, d := range data {
		sMsg := &QBFTSignedMessage{}
		if err := sMsg.UnmarshalSSZ(d); err != nil {
			return nil, err
		}
		ret[i] = sMsg
	}
	return ret, nil
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (msg *QBFTMessage) Validate() error {
	if _, err := msg.GetRoundChangeJustifications(); err != nil {
		return err
	}
	if _, err := msg.GetPrepareJustifications(); err != nil {
		return err
	}
	if msg.MsgType > RoundChangeMessageType {
		return errors.New("message type is invalid")
	}
	return nil
}

type QBFTSignedMessage struct {
	// Message is at the top for quick identifier look (see docs)
	Message   QBFTMessage
	Signature [96]byte `ssz-size:"96"`
	Signers   []uint64 `ssz-max:"13"`
	FullData  []byte   `ssz-max:"4259840"`
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (signedMsg *QBFTSignedMessage) Validate() error {
	if len(signedMsg.Signers) == 0 {
		return errors.New("message signers is empty")
	}

	// check unique signers
	signed := make(map[uint64]bool)
	for _, signer := range signedMsg.Signers {
		if signed[signer] {
			return errors.New("non unique signer")
		}
		if signer == 0 {
			return errors.New("signer ID 0 not allowed")
		}
		signed[signer] = true
	}

	return signedMsg.Message.Validate()
}

// WithoutFUllData returns SignedMessage without full data
func (signedMsg *QBFTSignedMessage) WithoutFUllData() *QBFTSignedMessage {
	return &QBFTSignedMessage{
		Signers:   signedMsg.Signers,
		Signature: signedMsg.Signature,
		Message:   signedMsg.Message,
	}
}
