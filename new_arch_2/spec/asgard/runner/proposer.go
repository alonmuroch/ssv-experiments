package runner

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func ProcessRandao(state *types.State, share *types.Share, message *types.SignedPartialSignatureMessages) error {
	panic("implement")
}

func ProposerExpectedPostConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	if decidedBlindedBlock(state) {
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

// decidedBlindedBlock returns true if decided value has a blinded block, false if regular block
// WARNING!! should be called after decided only
func decidedBlindedBlock(state *types.State) bool {
	_, _, err := DecidedConsensusData(state).GetBlindedBlockData()
	return err == nil
}
