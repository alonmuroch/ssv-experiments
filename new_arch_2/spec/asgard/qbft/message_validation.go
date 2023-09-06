package qbft

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// ValidateMessage returns nil if message valid for state
func ValidateMessage(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if err := signedMessage.Validate(); err != nil {
		return err
	}

	if signedMessage.Message.Round < state.Round {
		return errors.New("past round")
	}

	if signedMessage.Message.Height != state.Height {
		return errors.New("wrong message height")
	}

	switch signedMessage.Message.MsgType {
	case types.ProposalMessageType:
		return isValidProposal(state, share, signedMessage)
	case types.PrepareMessageType:
		if state.ProposalAcceptedForCurrentRound == nil {
			return errors.New("no proposal accepted for round")
		}
		return validSignedPrepareForHeightRoundAndRoot(share, signedMessage, state.Height, state.Round, state.ProposalAcceptedForCurrentRound.Message.Root)
	case types.CommitMessageType:
		return isValidCommit(state, share, signedMessage)
	case types.RoundChangeMessageType:
		// TODO validRoundChangeForData
		return nil
	default:
		return errors.New("unknown message type")
	}
}
