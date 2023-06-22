package qbft

type State struct {
	Round  uint64
	Height uint64

	PreparedRound uint64

	// Messages is a unified (to all message type) container slice, simple and easy to serialize
	Messages Container `ssz-max:"256"`
}

func (s *State) PrepareValue() *InputData {
	//prepare := s.Messages.RoundAndType(s.PreparedRound, PrepareMessageType)

	// TODO check quorum for certain value

	return &InputData{}
}

func (s *State) DecidedValue() *InputData {
	//prepare := s.Messages.RoundAndType(s.PreparedRound, CommitMessageType)

	// TODO check quorum for certain value

	return &InputData{}
}

func (s *State) AddMessage(msg *SignedMessage) {
	s.Messages = append(s.Messages, msg)
}
