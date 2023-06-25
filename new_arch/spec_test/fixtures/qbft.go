package fixtures

import (
	"ssv-experiments/new_arch/qbft"
)

func QBFTSignedMessage(signer, round, msgType uint64) *qbft.SignedMessage {
	return &qbft.SignedMessage{
		Message: qbft.Message{Round: round, MsgType: msgType},
		Signers: []uint64{signer},
	}
}
