package ssv

import (
	"ssv-experiments/new_arch/types"
)

type State struct {
	// PartialSignatures holds partial BLS signatures
	PartialSignatures Container `ssz-max:"256"`
	// DecidedValue holds the decided value set after consensus phase
	// TODO - Set as []byte because of SSZ limitations if it was ConsensusData (needs to be set for ssz to encode)
	DecidedValue []byte `ssz-max:"8388608"` // 2^23
	StartingDuty *types.Duty
}

func NewState(duty *types.Duty) *State {
	return &State{
		StartingDuty: duty,
	}
}

func (s *State) ReconstructSignature(msgs []*types.SignedPartialSignatureMessages) ([96]byte, error) {
	panic("implement")
}

// DecidedConsensusData will return decided consensus data or nil if not decided
func (s *State) DecidedConsensusData() *types.ConsensusData {
	panic("implement")
}
