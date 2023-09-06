package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponCommit(state *types.QBFT, msg *types.QBFTSignedMessage) error {
	// TOOD implement
	AddMessage(state, msg)
	return nil
}

func CreateCommitMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.CommitMessageType,
	}, nil
}

func CommitQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.CommitMessageType)
	return UniqueSignerQuorum(share.Quorum, all)
}
