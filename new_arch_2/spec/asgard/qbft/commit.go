package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func (i *Instance) UponCommit(msg *types.SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)
	return nil
}

func (i *Instance) CreateCommitMessage() (*types.Message, error) {
	// TODO implement
	return &types.Message{
		Round:   i.State.Round,
		MsgType: types.CommitMessageType,
	}, nil
}

func (i *Instance) CommitQuorum() bool {
	all := i.State.RoundAndType(i.State.Round, types.CommitMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
