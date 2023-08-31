package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func (i *Instance) UponRoundChange(msg *types.SignedMessage) error {
	panic("implement")
}

func (i *Instance) CreateRoundChangeMessage() (*types.SignedMessage, error) {
	panic("implement")
}

func (i *Instance) RoundChangeQuorum() bool {
	all := i.State.RoundAndType(i.State.Round, types.RoundChangeMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}

func (i *Instance) RoundChangePartialQuorum() bool {
	all := i.State.RoundAndType(i.State.Round, types.RoundChangeMessageType)
	if len(all) >= int(i.Share.PartialQuorum) {
		return true
	}
	return false
}
