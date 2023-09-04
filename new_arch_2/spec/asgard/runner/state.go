package runner

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
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

// ReconstructPostConsensusSignatures returns reconstructed signatures for post consensus messages
func ReconstructPostConsensusSignatures(state *types.State) ([]*bls.Sign, error) {
	ret := make([]*bls.Sign, 0)

	msgs := AllPostConsensus(state)

	for i := range msgs[0].Message.Signatures {
		signatures := make([][96]byte, 0)
		indexes := make([]uint64, 0)
		for _, m := range msgs {
			signatures = append(signatures, m.Message.Signatures[i].Signature)
			indexes = append(indexes, m.Signer)
		}

		reconstructedSig, err := reconstructSignatures(signatures, indexes)
		if err != nil {
			return nil, err
		}
		ret = append(ret, reconstructedSig)
	}
	return ret, nil
}

// reconstructSignatures takes array of signatures and indexes and returns a reconstructed bls.Sign object or error
func reconstructSignatures(signatures [][96]byte, indexes []uint64) (*bls.Sign, error) {
	reconstructedSig := bls.Sign{}

	idVec := make([]bls.ID, 0)
	sigVec := make([]bls.Sign, 0)

	for i, signature := range signatures {
		blsID := bls.ID{}
		err := blsID.SetDecString(fmt.Sprintf("%d", indexes[i]))
		if err != nil {
			return nil, err
		}

		idVec = append(idVec, blsID)
		blsSig := bls.Sign{}

		err = blsSig.Deserialize(signature[:])
		if err != nil {
			return nil, err
		}

		sigVec = append(sigVec, blsSig)
	}
	err := reconstructedSig.Recover(sigVec, idVec)
	return &reconstructedSig, err
}
