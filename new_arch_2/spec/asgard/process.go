package asgard

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/qbft"
	"ssv-experiments/new_arch_2/spec/asgard/runner"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// ProcessMessage processes a P2P message, modifying the state accordingly or returns error
func ProcessMessage(state *types.State, share *types.Share, message *types.Message) error {
	switch message.MsgType {
	case types.SSVConsensusMsgType:
		parsedMsg := &types.QBFTSignedMessage{}
		if err := parsedMsg.UnmarshalSSZ(message.Data); err != nil {
			return err
		}
		return qbft.ProcessMessage(state.QBFT, share, parsedMsg)
	case types.SSVPartialSignatureMsgType:
		parsedMsg := &types.SignedPartialSignatureMessages{}
		if err := parsedMsg.UnmarshalSSZ(message.Data); err != nil {
			return err
		}
		return runner.ProcessMessage(state, share, parsedMsg)

	default:
		return errors.New("unknown message type")
	}
}
