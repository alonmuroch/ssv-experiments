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
func (i *Instance) CreatePrepareMessage() (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   i.State.Round,
		MsgType: types.PrepareMessageType,
	}, nil
}

func (i *Instance) PrepareQuorum() bool {
	all := RoundAndType(i.State, i.State.Round, types.PrepareMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
