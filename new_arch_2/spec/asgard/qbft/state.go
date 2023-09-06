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

// uniqueSingerRound returns true if <signer, round, message type> unique
func uniqueSingerRound(state *types.QBFT, signedMessage *types.QBFTSignedMessage) bool {
	for _, msg := range state.Messages {
		uniqueSigners := func(a, b []uint64) bool {
			for i := range a {
				for j := range b {
					if a[i] == b[j] {
						return false
					}
				}
			}
			return true
		}(msg.Signers, signedMessage.Signers)

		if msg.Message.Round == state.Round &&
			msg.Message.MsgType == signedMessage.Message.MsgType &&
			!uniqueSigners {
			return false
		}
	}
	return true
}
