package fixtures

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

func QBFTSignedMessage(signer, round, msgType uint64) *qbft.SignedMessage {
	var fullData []byte
	root := [32]byte{}
	if msgType == qbft.ProposalMessageType {
		cd := AttesterConsensusData
		fullData, _ = cd.MarshalSSZ()
		root, _ = cd.HashTreeRoot()
	}

	return &qbft.SignedMessage{
		Message: qbft.Message{
			Round:      round,
			MsgType:    msgType,
			Root:       root,
			Identifier: p2p.NewIdentifier(Slot, ValidatorPK, types.BeaconRoleAttester),
		},
		Signers:  []uint64{signer},
		FullData: fullData,
	}
}
