package pipeline

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

// DecodeMessage decodes a P2P message, error if can't
func DecodeMessage(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
	msg, ok := objects[0].(*p2p.Message)
	if !ok {
		return errors.New("object not of type p2p.Message"), nil
	}

	switch msg.MsgType {
	case p2p.SSVPartialSignatureMsgType:
		signedMsg := &types.SignedPartialSignatureMessages{}
		if err := signedMsg.UnmarshalSSZ(msg.Data); err != nil {
			return errors.Wrap(err, "could not get partial signature Message from network Message"), nil
		}
		return nil, []interface{}{signedMsg}
	case p2p.SSVConsensusMsgType:
		signedMsg := &qbft.SignedMessage{}
		if err := signedMsg.UnmarshalSSZ(msg.Data); err != nil {
			return errors.Wrap(err, "could not get consensus Message from network Message"), nil
		}
		return nil, []interface{}{signedMsg}
	default:
		return errors.New("unsupported message type"), nil
	}
}

// Broadcast will broadcast a message to the SSV network
func Broadcast(t p2p.MsgType) func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		data := objects[0].(ssz.Marshaler)

		byts, err := data.MarshalSSZ()
		if err != nil {
			return err, nil
		}

		msg := &p2p.Message{
			MsgType: t,
			MsgID:   pipeline.Runner.Identifier,
			Data:    byts,
		}

		// broadcast

		return nil, []interface{}{msg}
	}
}

// BroadcastBeacon broadcasts to the beacon chain
func BroadcastBeacon(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
	for _, item := range objects {
		if _, encodable := item.(ssz.Marshaler); encodable {
			// broadcast
		}
	}
	return nil, nil
}
