package qbft

import (
	"github.com/pkg/errors"
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

func ProcessMessage(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if !CanProcessMessages(state) {
		return errors.New("can't process new messages")
	}

	if err := ValidateMessage(state, share, signedMessage); err != nil {
		return err
	}

	switch signedMessage.Message.MsgType {
	case types.ProposalMessageType:
		return UponProposal(state, signedMessage)
	case types.PrepareMessageType:
		return UponPrepare(state, share, signedMessage)
	case types.CommitMessageType:
		return UponCommit(state, share, signedMessage)
	case types.RoundChangeMessageType:
		// TODO validRoundChangeForData
		return nil
	default:
		return errors.New("unknown signedMessage type")
	}
}

// CanProcessMessages returns true if can process messages
func CanProcessMessages(state *types.QBFT) bool {
	return !state.Stopped && int(state.Round) < CutoffRound
}

// IsProposer returns true if proposer for current round
func IsProposer(state *types.QBFT, share *types.Share) bool {
	return proposerForRound(state.Round) == share.OperatorID
}

func proposerForRound(round uint64) uint64 {
	// TODO round robin
	return 1
}

// DecidedRoot returns the root and messages that decided current round
func DecidedRoot(state *types.QBFT, share *types.Share) (bool, []*types.QBFTSignedMessage, [32]byte) {
	byRoot := make(map[[32]byte][]*types.QBFTSignedMessage)

	// batch messages by root. If exists a decided message return immediately
	for _, m := range RoundAndType(state, state.Round, types.CommitMessageType) {
		// decided message return true
		if len(m.Signers) >= int(share.Quorum) {
			return true, []*types.QBFTSignedMessage{m}, m.Message.Root
		}

		if byRoot[m.Message.Root] == nil {
			byRoot[m.Message.Root] = []*types.QBFTSignedMessage{}
		}
		byRoot[m.Message.Root] = append(byRoot[m.Message.Root], m)
	}

	// find if decided
	for _, msgs := range byRoot {
		if len(msgs) >= int(share.Quorum) {
			return true, msgs, msgs[0].Message.Root
		}
	}
	return false, nil, [32]byte{}
}

// DecidedValue returns decided value for current round
func DecidedValue(state *types.QBFT, share *types.Share) ([]byte, error) {
	decided, msgs, _ := DecidedRoot(state, share)
	if !decided {
		return nil, errors.New("not decided")
	}

	// single decided message
	if len(msgs) == 1 {
		return msgs[0].FullData, nil
	}

	// regular commit quorum
	proposalMsgs := RoundAndType(state, state.Round, types.ProposalMessageType)
	if len(proposalMsgs) != 1 {
		return nil, errors.New("no valid proposal for round")
	}
	return proposalMsgs[0].FullData, nil
}

func HashDataRoot(data []byte) ([32]byte, error) {
	return types.SSZBytes(data).HashTreeRoot()
}

// UniqueSignerQuorum returns true if all messages have 1 unique signer
func UniqueSignerQuorum(quorum uint64, messages []*types.QBFTSignedMessage) bool {
	uniqueSigners := make(map[uint64]bool)
	for _, msg := range messages {
		if len(msg.Signers) != 1 {
			return false
		}
		if uniqueSigners[msg.Signers[0]] {
			return false
		}
		uniqueSigners[msg.Signers[0]] = true
	}
	return uint64(len(uniqueSigners)) > quorum
}
