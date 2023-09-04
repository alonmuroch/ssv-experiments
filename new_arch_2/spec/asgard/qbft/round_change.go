package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponRoundChange(msg *types.QBFTSignedMessage) error {
	panic("implement")
}

func CreateRoundChangeMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	panic("implement")
}

func RoundChangeQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.RoundChangeMessageType)
	if len(all) >= int(share.Quorum) {
		return true
	}
	return false
}

func RoundChangePartialQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.RoundChangeMessageType)
	if len(all) >= int(share.PartialQuorum) {
		return true
	}
	return false
}
