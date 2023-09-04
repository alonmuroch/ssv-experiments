package runner

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func AttesterExpectedPostConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	attestationData, err := DecidedConsensusData(state).GetAttestationData()
	if err != nil {
		return nil, errors.Wrap(err, "could not get attestation data")
	}
	return []ssz.HashRoot{attestationData}, nil
}
