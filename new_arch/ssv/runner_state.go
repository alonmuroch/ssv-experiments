package ssv

import "ssv-experiments/new_arch/types"

type State struct {
	// PartialSignatures holds partial BLS signatures
	PartialSignatures *PartialSignatureContainer

	StartingDuty *types.Duty
	DecidedData  *types.ConsensusData
}

func NewState(duty *types.Duty) State {
	return State{
		StartingDuty: duty,
	}
}
