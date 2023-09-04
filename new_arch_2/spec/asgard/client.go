package asgard

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/qbft"
	"ssv-experiments/new_arch_2/spec/asgard/runner"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// StartAttesterRunner starts an attester runner with QBFT instance
func StartAttesterRunner(state *types.State, share *types.Share, consensusData *types.ConsensusData) (interface{}, error) {
	state.StartingDuty = consensusData.Duty

	return qbft.Start(
		state.QBFT,
		share,
		consensusData,
		consensusData.Duty.Slot,
	)
}

// StartProposerRunner starts aproposer runner with pre-consensus
func StartProposerRunner(state *types.State, share *types.Share, duty *types.Duty, randaoRoot [32]byte, signedRandao [96]byte) (interface{}, error) {
	state.StartingDuty = duty

	return &types.PartialSignatureMessages{
		Type: types.BeaconRoleProposer,
		Slot: duty.Slot,
		Signatures: []*types.PartialSignatureMessage{
			{
				Signature: signedRandao,
				Root:      randaoRoot,
			},
		},
	}, nil
}

// UponPreConsensusQuorum starts a QBFT instance, returns a QBFT message or error
func UponPreConsensusQuorum(state *types.State, share *types.Share, startingConsensusData *types.ConsensusData) (*types.QBFTMessage, error) {
	return qbft.Start(
		state.QBFT,
		share,
		startingConsensusData,
		startingConsensusData.Duty.Slot,
	)
}

// UponPostConsensusQuorum returns reconstructed signed object to broadcast, or error
func UponPostConsensusQuorum(state *types.State) (interface{}, error) {
	role := runner.Role(state)
	switch role {
	case types.BeaconRoleAttester:
		return runner.ReconstructAttestationData(state)
	case types.BeaconRoleProposer:
		if runner.DecidedBlindedBlock(state) {
			return runner.ReconstructBlindedBlock(state)
		}
		return runner.ReconstructBlock(state)
	default:
		return nil, errors.New("role not supported")
	}
}
