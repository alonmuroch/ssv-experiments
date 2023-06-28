package qbft

func (i *Instance) UponRoundChange(msg *SignedMessage) error {
	panic("implement")
}

func (i *Instance) CreateRoundChangeMessage() (*SignedMessage, error) {
	panic("implement")
}

func (i *Instance) RoundChangeQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, RoundChangeMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}

func (i *Instance) RoundChangePartialQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, RoundChangeMessageType)
	if len(all) >= int(i.Share.PartialQuorum) {
		return true
	}
	return false
}
