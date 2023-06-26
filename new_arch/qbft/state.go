package qbft

import "ssv-experiments/new_arch/types"

type State struct {
	Round  uint64
	Height uint64

	PreparedRound uint64

	// Messages is a unified (to all message type) container slice, simple and easy to serialize.
	// All messages in the container are verified and authenticated
	Messages Container `ssz-max:"256"`
}

func (s *State) PrepareValue() *types.ConsensusData {
	//prepare := s.Messages.RoundAndType(s.PreparedRound, PrepareMessageType)

	// TODO check quorum for certain value

	return &types.ConsensusData{}
}

func (s *State) AddMessage(msg *SignedMessage) {
	s.Messages = append(s.Messages, msg)
}
