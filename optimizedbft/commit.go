package optimizedbft

func (s *State) ProcessCommit(msg *Message) error {
	if err := s.BaseMsgValidation(msg, Commit); err != nil {
		return err
	}

	s.AddMessage(msg)
	return nil
}
