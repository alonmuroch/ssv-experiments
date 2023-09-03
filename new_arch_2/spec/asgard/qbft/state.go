package qbft

import (
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

func RoundAndType(state *types.QBFT, round uint64, msgType uint64) []*types.QBFTSignedMessage {
	ret := make([]*types.QBFTSignedMessage, 0)
	for _, msg := range state.Messages {
		if msg.Message.Round == round && msg.Message.MsgType == msgType {
			ret = append(ret, msg)
		}
	}
	return ret
}

func AddMessage(state *types.QBFT, msg *types.QBFTSignedMessage) {
	state.Messages = append(state.Messages, msg)
}
