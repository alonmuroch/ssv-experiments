package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponProposal(state *types.QBFT, msg *types.QBFTSignedMessage) error {
	// TOOD implement
	AddMessage(state, msg)
	return nil
}

func CreateProposalMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.ProposalMessageType,
	}, nil
}
