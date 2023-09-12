package runner

import (
	"github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec"
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponProposerDecided returns an array of partial sig messages to be signed and broadcasted
func UponProposerDecided(state *types.State) ([]*types.PartialSignatureMessage, error) {
	cd := DecidedConsensusData(state)

	var blkToSign ssz.HashRoot
	var err error
	if DecidedBlindedBlock(state) {
		_, blkToSign, err = cd.GetBlindedBlockData()
	} else {
		_, blkToSign, err = cd.GetBlockData()
	}
	if err != nil {
		return nil, err
	}

	domainData := cd.Duty.DomainData
	root, err := types.ComputeETHSigningRoot(blkToSign, domainData)
	if err != nil {
		return nil, err
	}

	return []*types.PartialSignatureMessage{
		{
			Root: root,
		},
	}, nil
}

// ReconstructBlock returns an object with the versioned block and reconstructed signature or error
func ReconstructBlock(state *types.State) (interface{}, error) {
	cd := DecidedConsensusData(state)
	if cd == nil {
		return nil, errors.New("consensus data nil")
	}
	blk, _, err := cd.GetBlockData()
	if err != nil {
		return nil, err
	}

	sigs, err := ReconstructPostConsensusSignatures(state)
	if err != nil {
		return nil, err
	}

	return struct {
		Block     *spec.VersionedBeaconBlock
		Signature [96]byte
	}{
		Block:     blk,
		Signature: types.SignToBLSSignature(sigs[0]),
	}, nil
}

func ReconstructBlindedBlock(state *types.State) (interface{}, error) {
	cd := DecidedConsensusData(state)
	if cd == nil {
		return nil, errors.New("consensus data nil")
	}
	blk, _, err := cd.GetBlindedBlockData()
	if err != nil {
		return nil, err
	}

	sigs, err := ReconstructPostConsensusSignatures(state)
	if err != nil {
		return nil, err
	}

	return struct {
		VersionedBlock *api.VersionedBlindedBeaconBlock
		Signature      [96]byte
	}{
		VersionedBlock: blk,
		Signature:      types.SignToBLSSignature(sigs[0]),
	}, nil
}

func ProcessRandao(state *types.State, share *types.Share, message *types.SignedPartialSignatureMessages) error {
	panic("implement")
}

func ProposerExpectedPreConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	epoch := state.StartingDuty.BeaconNetwork.EstimatedEpochAtSlot(state.StartingDuty.Slot)
	return []ssz.HashRoot{types.SSZUint64(epoch)}, nil
}

func ProposerExpectedPostConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	if DecidedBlindedBlock(state) {
		_, data, err := DecidedConsensusData(state).GetBlindedBlockData()
		if err != nil {
			return nil, errors.Wrap(err, "could not get blinded block data")
		}
		return []ssz.HashRoot{data}, nil
	}

	_, data, err := DecidedConsensusData(state).GetBlockData()
	if err != nil {
		return nil, errors.Wrap(err, "could not get block data")
	}
	return []ssz.HashRoot{data}, nil
}

// DecidedBlindedBlock returns true if decided value has a blinded block, false if regular block
// WARNING!! should be called after decided only
func DecidedBlindedBlock(state *types.State) bool {
	_, _, err := DecidedConsensusData(state).GetBlindedBlockData()
	return err == nil
}

// ProposerValidateConsensusData validates consensus data, returns nil if valid
func ProposerValidateConsensusData(cd *types.ConsensusData) error {
	panic("implement")
}
