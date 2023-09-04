package runner

import (
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

// Role returns the state's beacon role
func Role(state *types.State) uint64 {
	return state.StartingDuty.Role
}

func AllPreConsensus(state *types.State) []*types.SignedPartialSignatureMessages {
	ret := make([]*types.SignedPartialSignatureMessages, 0)
	for _, m := range state.PartialSignatures {
		if m.Message.Type.IsPreConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}

func AllPostConsensus(state *types.State) []*types.SignedPartialSignatureMessages {
	ret := make([]*types.SignedPartialSignatureMessages, 0)
	for _, m := range state.PartialSignatures {
		if m.Message.Type.IsPostConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}

// DecidedConsensusData will return decided consensus data or nil if not decided
func DecidedConsensusData(state *types.State) *types.ConsensusData {
	if state.QBFT == nil {
		return nil
	}
	return state.QBFT.DecidedValue()
}
