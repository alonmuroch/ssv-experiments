package qbft

func (i *Instance) UponPrepare(msg *SignedMessage) error {
	panic("implement")
}

func (i *Instance) CreatePrepareMessage() (*SignedMessage, error) {
	panic("implement")
}

func (i *Instance) PrepareQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, PrepareMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
