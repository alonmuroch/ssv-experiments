package fixtures

import "ssv-experiments/new_arch/types"

var AttesterConsensusData = &types.ConsensusData{
	Duty:        AttesterDuty,
	DataVersion: 0,
	DataSSZ:     AttestationDataBytes,
}
