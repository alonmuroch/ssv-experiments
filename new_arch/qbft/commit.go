package qbft

func (i *Instance) UponCommit(msg *SignedMessage) error {
	panic("implement")
}

func (i *Instance) CreateCommitMessage() (*SignedMessage, error) {
	panic("implement")
}

func (i *Instance) CommitQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, CommitMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
