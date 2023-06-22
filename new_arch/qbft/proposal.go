package qbft

func (i *Instance) UponProposal(msg *SignedMessage) error {
	// TOOD implement
	i.State.AddMessage(msg)
	return nil
}

func (i *Instance) CreateProposalMessage() (*Message, error) {
	// TODO implement
	return &Message{
		Round:   i.State.Round,
		MsgType: ProposalMessageType,
	}, nil
}
