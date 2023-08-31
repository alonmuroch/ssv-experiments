package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func (i *Instance) UponProposal(msg *types.SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)
	return nil
}

func (i *Instance) CreateProposalMessage() (*types.Message, error) {
	// TODO implement
	return &types.Message{
		Round:   i.State.Round,
		MsgType: types.ProposalMessageType,
	}, nil
}
