package runner

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/go-bitfield"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponAttesterDecided returns an array of partial sig messages to be signed and broadcasted
func UponAttesterDecided(state *types.State) ([]*types.PartialSignatureMessage, error) {
	cd := DecidedConsensusData(state)
	attData, err := cd.GetAttestationData()
	if err != nil {
		return nil, err
	}

	domainData := cd.Duty.DomainData
	root, err := types.ComputeETHSigningRoot(attData, domainData)
	if err != nil {
		return nil, err
	}

	return []*types.PartialSignatureMessage{
		{
			Root: root,
		},
	}, nil
}

func ReconstructAttestationData(state *types.State) (*phase0.Attestation, error) {
	cd := DecidedConsensusData(state)
	if cd == nil {
		return nil, errors.New("consensus data nil")
	}
	attData, err := cd.GetAttestationData()
	if err != nil {
		return nil, err
	}

	sigs, err := ReconstructPostConsensusSignatures(state)
	if err != nil {
		return nil, err
	}

	aggregationBitfield := bitfield.NewBitlist(cd.Duty.CommitteeLength)
	aggregationBitfield.SetBitAt(cd.Duty.ValidatorCommitteeIndex, true)
	ret := &phase0.Attestation{
		Data:            attData,
		Signature:       types.SignToBLSSignature(sigs[0]),
		AggregationBits: aggregationBitfield,
	}
	return ret, nil
}

func AttesterExpectedPostConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	attestationData, err := DecidedConsensusData(state).GetAttestationData()
	if err != nil {
		return nil, errors.Wrap(err, "could not get attestation data")
	}
	return []ssz.HashRoot{attestationData}, nil
}

// AttesterValidateConsensusData validates consensus data, returns nil if valid
func AttesterValidateConsensusData(cd *types.ConsensusData) error {
	panic("implement")
}
