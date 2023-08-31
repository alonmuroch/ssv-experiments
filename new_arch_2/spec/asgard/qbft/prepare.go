package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func (i *Instance) UponPrepare(msg *types.SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)

	prepareMsg := i.State.RoundAndType(i.State.Round, types.PrepareMessageType)
	if len(prepareMsg) >= int(i.Share.Quorum) {
		i.State.PreparedRound = i.State.Round
	}

	return nil
}

// CreatePrepareMessage returns unsigned prepare message
func (i *Instance) CreatePrepareMessage() (*types.Message, error) {
	// TODO implement
	return &types.Message{
		Round:   i.State.Round,
		MsgType: types.PrepareMessageType,
	}, nil
}

func (i *Instance) PrepareQuorum() bool {
	all := i.State.RoundAndType(i.State.Round, types.PrepareMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
