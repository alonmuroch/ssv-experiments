package ssv

import (
	"ssv-experiments/new_arch/types"
)

type State struct {
	Share *types.Share
	// PartialSignatures holds partial BLS signatures
	PartialSignatures Container

	StartingDuty *types.Duty
	DecidedData  *types.ConsensusData
}

func NewState(duty *types.Duty) State {
	return State{
		StartingDuty: duty,
	}
}

func (s *State) HasPreConsensusQuorum() bool {
	all := s.PartialSignatures.AllPreConsensus()
	return len(all) >= int(s.Share.Quorum)
}

func (s *State) HasPostConsensusQuorum() bool {
	all := s.PartialSignatures.AllPostConsensus()
	return len(all) >= int(s.Share.Quorum)
}

func (s *State) ReconstructSignature(msgs []*types.SignedPartialSignatureMessages) ([96]byte, error) {
	panic("implement")
}
