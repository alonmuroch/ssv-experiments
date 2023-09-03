package qbft

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func ValidateMessage(state *types.QBFT, message *types.QBFTSignedMessage) error {
	if err := message.Validate(); err != nil {
		return err
	}

	if message.Message.Round < state.Round {
		return errors.New("past round")
	}

	switch message.Message.MsgType {
	case types.ProposalMessageType:
		// TODO isValidProposal
		return nil
	case types.PrepareMessageType:
		// TODO validSignedPrepareForHeightRoundAndRoot
		return nil
	case types.CommitMessageType:
		// TODO validateCommit
		return nil
	case types.RoundChangeMessageType:
		// TODO validRoundChangeForData
		return nil
	default:
		return errors.New("unknown message type")
	}
}
