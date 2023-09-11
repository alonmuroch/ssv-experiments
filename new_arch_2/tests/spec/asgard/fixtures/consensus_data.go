package fixtures

import "ssv-experiments/new_arch_2/spec/asgard/types"

var AttesterConsensusData = &types.ConsensusData{
	Duty:        AttesterDuty,
	DataVersion: 0,
	DataSSZ:     AttestationDataBytes,
}

var ProposerConsensusData = &types.ConsensusData{
	Duty:        ProposerDuty,
	DataVersion: 0,
	DataSSZ:     BeaconBlockCapellaBytes,
}
