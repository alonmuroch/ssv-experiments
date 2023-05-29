package pipeline

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/ssv"
)

func Broadcast(t p2p.MsgType) func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		data := objects[0].(ssz.Marshaler)

		byts, err := data.MarshalSSZ()
		if err != nil {
			return err, nil
		}

		msg := &p2p.Message{
			MsgType: t,
			MsgID:   runner.Identifier,
			Data:    byts,
		}

		// broadcast

		return nil, []interface{}{msg}
	}
}

// BroadcastBeacon broadcasts to the beacon chain
func BroadcastBeacon(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	for _, item := range objects {
		if _, encodable := item.(ssz.Marshaler); encodable {
			// broadcast
		}
	}
	return nil, nil
}
