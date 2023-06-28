package qbft

func (i *Instance) UponCommit(msg *SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)
	return nil
}

func (i *Instance) CreateCommitMessage() (*Message, error) {
	// TODO implement
	return &Message{
		Round:   i.State.Round,
		MsgType: CommitMessageType,
	}, nil
}

func (i *Instance) CommitQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, CommitMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
