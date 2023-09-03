package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponCommit(state *types.QBFT, msg *types.QBFTSignedMessage) error {
	// TOOD implement
	AddMessage(state, msg)
	return nil
}

func (i *Instance) CreateCommitMessage() (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   i.State.Round,
		MsgType: types.CommitMessageType,
	}, nil
}

func (i *Instance) CommitQuorum() bool {
	all := RoundAndType(i.State, i.State.Round, types.CommitMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
