package qbft

func (i *Instance) UponPrepare(msg *SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)

	prepareMsg := i.State.Messages.RoundAndType(i.State.Round, PrepareMessageType)
	if len(prepareMsg) >= int(i.Share.Quorum) {
		i.State.PreparedRound = i.State.Round
	}

	return nil
}

// CreatePrepareMessage returns unsigned prepare message
func (i *Instance) CreatePrepareMessage() (*Message, error) {
	// TODO implement
	return &Message{
		Round:   i.State.Round,
		MsgType: PrepareMessageType,
	}, nil
}

func (i *Instance) PrepareQuorum() bool {
	all := i.State.Messages.RoundAndType(i.State.Round, PrepareMessageType)
	if len(all) >= int(i.Share.Quorum) {
		return true
	}
	return false
}
