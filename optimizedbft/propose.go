package optimizedbft

import (
	"bytes"
	"fmt"
)

func (s *State) ProcessPropose(share *Share, msg *Message) error {
	if err := s.BaseMsgValidation(msg, Propose); err != nil {
		return err
	}

	if s.RoundProposer() != msg.Round {
		return fmt.Errorf("wrong proposer")
	}

	if s.PreparedRound != 0 {
		data, err := s.MessageQuorumDataForRound(Prepare, s.PreparedRound)
		if err != nil {
			return err
		}
		if !bytes.Equal(msg.Data, data) {
			return fmt.Errorf("propose data invalid")
		}
	}

	s.AddMessage(msg)
	return nil
}
