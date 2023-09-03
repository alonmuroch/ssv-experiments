package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponProposal(state *types.QBFT, msg *types.QBFTSignedMessage) error {
	// TOOD implement
	AddMessage(state, msg)
	return nil
}

func (i *Instance) CreateProposalMessage() (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   i.State.Round,
		MsgType: types.ProposalMessageType,
	}, nil
}
