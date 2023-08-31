package runner

import (
	types2 "ssv-experiments/new_arch_2/spec/asgard/types"
)

type State types2.State

func NewState(duty *types2.Duty) *State {
	return &State{
		StartingDuty:      duty,
		PartialSignatures: []*types2.SignedPartialSignatureMessages{},
	}
}

func (state *State) AllPreConsensus() []*types2.SignedPartialSignatureMessages {
	ret := make([]*types2.SignedPartialSignatureMessages, 0)
	for _, m := range state.PartialSignatures {
		if m.Message.Type.IsPreConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}

func (state *State) AllPostConsensus() []*types2.SignedPartialSignatureMessages {
	ret := make([]*types2.SignedPartialSignatureMessages, 0)
	for _, m := range state.PartialSignatures {
		if m.Message.Type.IsPostConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}

// DecidedConsensusData will return decided consensus data or nil if not decided
func (state *State) DecidedConsensusData() *types2.ConsensusData {
	if state.QBFT == nil {
		return nil
	}
	return state.QBFT.DecidedValue()
}
