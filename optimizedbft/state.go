package optimizedbft

import "fmt"

type State struct {
	Round                           uint64
	Height                          uint64
	PreparedRound                   uint64
	ProposalAcceptedForCurrentRound *Message
	Messages                        []*Message
	Stopped                         bool
}

func NewState() *State {
	return &State{
		Round:    1,
		Messages: make([]*Message, 0),
	}
}

func (s *State) BumpRound() {
	s.Round++
	s.ProposalAcceptedForCurrentRound = nil
}

func (s *State) RoundProposer() uint64 {
	return s.Round
}

func (s *State) AddMessage(msg *Message) {
	s.Messages = append(s.Messages, msg)
}

func (s *State) MessageQuorumForRound(share *Share, msgType MessageType, round uint64) bool {
	signers := map[uint64]bool{}
	for _, msg := range s.Messages {
		if msg.Round == round && msg.Type == msgType && !signers[msg.OperatorID] {
			signers[msg.OperatorID] = true
		}
	}

	return uint64(len(signers)) >= share.Quorum
}

func (s *State) MessageQuorumDataForRound(msgType MessageType, round uint64) ([]byte, error) {
	for _, msg := range s.Messages {
		if msg.Round == round && msg.Type == msgType {
			return msg.Data, nil
		}
	}
	return nil, fmt.Errorf("data not found")
}

func (s *State) BaseMsgValidation(msg *Message, msgType MessageType) error {
	if s.Stopped {
		return fmt.Errorf("stopped")
	}

	if s.Round != msg.Round {
		return fmt.Errorf("wrong round")
	}

	if s.Height != msg.Height {
		return fmt.Errorf("wrong height")
	}

	if msgType != msg.Type {
		return fmt.Errorf("wrong msg type")
	}

	return nil
}
