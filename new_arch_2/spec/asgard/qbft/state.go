package qbft

import (
	types2 "ssv-experiments/new_arch_2/spec/asgard/types"
)

type State types2.QBFT

func (state *State) RoundAndType(round uint64, msgType uint64) []*types2.SignedMessage {
	ret := make([]*types2.SignedMessage, 0)
	for _, msg := range state.Messages {
		if msg.Message.Round == round && msg.Message.MsgType == msgType {
			ret = append(ret, msg)
		}
	}
	return ret
}

func (s *State) AddMessage(msg *types2.SignedMessage) {
	s.Messages = append(s.Messages, msg)
}
