package qbft

type Container []*SignedMessage

func (c Container) All() *SignedMessage {
	panic("implement")
}

func (c Container) Round(round uint64) []*SignedMessage {
	panic("implement")
}

func (c Container) RoundAndRoot(round uint64, root [32]byte) []*SignedMessage {
	panic("implement")
}
