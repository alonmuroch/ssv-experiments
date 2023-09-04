package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponPrepare(state *types.QBFT, share *types.Share, msg *types.QBFTSignedMessage) error {
	// TOOD implement
	AddMessage(state, msg)

	prepareMsg := RoundAndType(state, state.Round, types.PrepareMessageType)
	if len(prepareMsg) >= int(share.Quorum) {
		state.PreparedRound = state.Round
	}

	return nil
}

// CreatePrepareMessage returns unsigned prepare message
func CreatePrepareMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.PrepareMessageType,
	}, nil
}

func PrepareQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.PrepareMessageType)
	if len(all) >= int(share.Quorum) {
		return true
	}
	return false
}
