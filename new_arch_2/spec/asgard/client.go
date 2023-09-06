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

// UponProposal returns a prepare message for a valid proposal
func UponProposal(state *types.State) (*types.QBFTMessage, error) {
	if !qbft.CanProcessMessages(state.QBFT) {
		return nil, errors.New("can't process events/ messages")
	}

	// TODO - i.config.GetTimer().TimeoutForRound(signedProposal.Message.Round)

	if err := validateConsensusData(state, state.QBFT.ProposalAcceptedForCurrentRound.FullData); err != nil {
		return nil, err
	}

	msg, err := qbft.CreatePrepareMessage(state.QBFT)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// UponPrepareQuorum returns a commit message upon prepare quorum
func UponPrepareQuorum(state *types.QBFT) (*types.QBFTMessage, error) {
	if !qbft.CanProcessMessages(state) {
		return nil, errors.New("can't process events/ messages")
	}

	msg, err := qbft.CreateCommitMessage(state)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// UponCommitQuorum moves to post consensus phase, returns an array of partial signature messages to sign and broadcast
func UponCommitQuorum(state *types.State, share *types.Share) ([]*types.PartialSignatureMessage, error) {
	byts, _ := qbft.DecidedValue(state.QBFT, share) // no need to handle error, validateConsensusData will handle it
	if err := validateConsensusData(state, byts); err != nil {
		return nil, err
	}

	role := runner.Role(state)
	switch role {
	case types.BeaconRoleAttester:
		return runner.UponAttesterDecided(state)
	case types.BeaconRoleProposer:
		return runner.UponProposerDecided(state)
	default:
		return nil, errors.New("role not supported")
	}
}

// UponTimeout bumps round and returns round change for timer timeout
func UponTimeout(state *types.QBFT) (*types.QBFTMessage, error) {
	if !qbft.CanProcessMessages(state) {
		return nil, errors.New("can't process events/ messages")
	}

	newRound := state.Round + 1
	defer func() {
		state.Round = newRound
		state.ProposalAcceptedForCurrentRound = nil
	}()

	msg, err := qbft.CreateRoundChangeMessage(state)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// UponRoundChangeQuorum broadcasts a proposal, if proposer, message upon round change quorum
func UponRoundChangeQuorum(state *types.QBFT) error {
	panic("implement")
}

// UponF1RoundChangeQuorum bumps round and broadcasts round change upon f+1 round changes
func UponF1RoundChangeQuorum(state *types.QBFT) error {
	panic("implement")
}

func validateConsensusData(state *types.State, data []byte) error {
	cd := &types.ConsensusData{}
	if err := cd.UnmarshalSSZ(data); err != nil {
		return err
	}

	role := runner.Role(state)
	switch role {
	case types.BeaconRoleAttester:
		return runner.AttesterValidateConsensusData(cd)
	case types.BeaconRoleProposer:
		return runner.ProposerValidateConsensusData(cd)
	default:
		return errors.New("role not supported")
	}
}
