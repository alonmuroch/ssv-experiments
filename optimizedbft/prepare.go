package optimizedbft

func (s *State) ProcessPrepare(share *Share, msg *Message) error {
	if err := s.BaseMsgValidation(msg, Prepare); err != nil {
		return err
	}

	s.AddMessage(msg)
	if s.MessageQuorumForRound(share, Prepare, s.Round) {
		s.PreparedRound = s.Round
	}
	return nil
}
