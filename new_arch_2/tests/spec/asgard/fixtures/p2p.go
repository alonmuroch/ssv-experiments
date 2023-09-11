package fixtures

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
	types2 "ssv-experiments/new_arch_2/spec/asgard/types"
)

var Identifier = p2p.NewIdentifier(Slot, ValidatorPK, types.BeaconRoleAttester)

func P2PConsensusMessage(round, signer, msgType uint64) *p2p.Message {
	m := &qbft.SignedMessage{
		Message: qbft.Message{Round: 1, MsgType: qbft.ProposalMessageType},
		Signers: []uint64{1},
	}

	byts, err := m.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}

	return &p2p.Message{
		MsgType: p2p.SSVConsensusMsgType,
		Data:    byts,
	}
}

func P2PPartialSignatureMessage(signer, slot uint64, msgType types2.PartialSigMsgType) *p2p.Message {
	m := SignedPartialSignatureMessage(signer, slot, msgType)

	byts, err := m.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}

	return &p2p.Message{
		Identifier: Identifier,
		MsgType:    p2p.SSVPartialSignatureMsgType,
		Data:       byts,
	}
}

func P2PQBFTSignedMessage(signer, round, msgType uint64) *p2p.Message {
	m := QBFTSignedMessage(signer, round, msgType)
	byts, err := m.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}

	return &p2p.Message{
		Identifier: Identifier,
		MsgType:    p2p.SSVConsensusMsgType,
		Data:       byts,
	}
}
