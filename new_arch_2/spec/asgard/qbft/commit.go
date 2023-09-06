package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponCommit returns true if a quorum of commit messages was received.
// Assumes commit message is valid!
func UponCommit(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if !uniqueSingerRound(state, signedMessage) {
		return errors.New("duplicate message")
	}
	AddMessage(state, signedMessage)

	return nil
}

func CreateCommitMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.CommitMessageType,
	}, nil
}

func CommitQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.CommitMessageType)
	return UniqueSignerQuorum(share.Quorum, all)
}

// isValidCommit returns nil if commit message (not a decided message) is valid for state
func isValidCommit(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if err := baseCommitValidation(share, state.Height, signedMessage); err != nil {
		return err
	}

	if state.ProposalAcceptedForCurrentRound == nil {
		return errors.New("no proposal accepted for round")
	}

	if len(signedMessage.Signers) != 1 {
		return errors.New("msg allows 1 signer")
	}

	if signedMessage.Message.Round != state.Round {
		return errors.New("wrong msg round")
	}

	if !bytes.Equal(state.ProposalAcceptedForCurrentRound.Message.Root[:], signedMessage.Message.Root[:]) {
		return errors.New("proposed data mismatch")
	}

	return nil
}

// baseCommitValidation returns true if commit message (which can be a decided message as well) is valid
func baseCommitValidation(share *types.Share, height uint64, signedMessage *types.QBFTSignedMessage) error {
	if signedMessage.Message.MsgType != types.CommitMessageType {
		return errors.New("commit msg type is wrong")
	}
	if signedMessage.Message.Height != height {
		return errors.New("wrong msg height")
	}

	if err := signedMessage.Validate(); err != nil {
		return errors.Wrap(err, "signed commit invalid")
	}

	// verify signature
	if err := types.VerifyObjectSignature(
		signedMessage.Signature,
		signedMessage,
		share.Domain,
		types.QBFTSignatureType,
		share.Cluster,
	); err != nil {
		return err
	}

	return nil
}
