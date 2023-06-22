package qbft

type Container []*SignedMessage

func NewContainer() Container {
	return []*SignedMessage{}
}

func (c Container) All() *SignedMessage {
	panic("implement")
}

func (c Container) Round(round uint64) []*SignedMessage {
	panic("implement")
}

func (c Container) RoundAndRoot(round uint64, root [32]byte) []*SignedMessage {
	panic("implement")
}

func (c Container) RoundAndType(round uint64, msgType uint64) []*SignedMessage {
	ret := make([]*SignedMessage, 0)
	for _, msg := range c {
		if msg.Message.Round == round && msg.Message.MsgType == msgType {
			ret = append(ret, msg)
		}
	}
	return ret
}
