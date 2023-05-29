package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

// ContinueIfConsensusMessage execute the next item if first object is a QBFT Signed Message, otherwise skip next
func ContinueIfConsensusMessage(next PipelineF) PipelineF {
	return func(runner *Runner, objects ...interface{}) (error, []interface{}) {
		if _, isConsensusMessage := objects[0].(*qbft.SignedMessage); isConsensusMessage {
			return next(runner, objects)
		}
		return nil, []interface{}{SkipNext}
	}
}

// ContinueIfPostConsensusMessage execute the next item if first object is a post consensus Message, otherwise skip next
func ContinueIfPostConsensusMessage(next PipelineF) PipelineF {
	return func(runner *Runner, objects ...interface{}) (error, []interface{}) {
		if m, isPartialSigMessage := objects[0].(*types.SignedPartialSignatureMessages); isPartialSigMessage && m.Message.Type == types.PostConsensusPartialSig {
			return next(runner, objects)
		}
		return nil, []interface{}{SkipNext}
	}
}

// ValidateDecidedValue returns a pipeline function for a specific value check function
func ValidateDecidedValue(valueCheck func(data *types.ConsensusData) error) PipelineF {
	return func(runner *Runner, objects ...interface{}) (error, []interface{}) {
		if err := valueCheck(objects[0].(*types.ConsensusData)); err != nil {
			return err, nil
		}
		return nil, objects
	}
}

func Broadcast(t p2p.MsgType) func(runner *Runner, objects ...interface{}) (error, []interface{}) {
	return func(runner *Runner, objects ...interface{}) (error, []interface{}) {
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
func BroadcastBeacon(runner *Runner, objects ...interface{}) (error, []interface{}) {
	for _, item := range objects {
		if _, encodable := item.(ssz.Marshaler); encodable {
			// broadcast
		}
	}
	return nil, nil
}

// ConstructPostConsensusMessage receives consensus data and partial sig message and returns PartialSignatureMessages
func ConstructPostConsensusMessage(t types.PartialSigMsgType) func(runner *Runner, objects ...interface{}) (error, []interface{}) {
	return func(runner *Runner, objects ...interface{}) (error, []interface{}) {
		cd := objects[0].(*types.ConsensusData)
		m := objects[1].(*types.PartialSignatureMessage)
		return nil, []interface{}{
			&types.PartialSignatureMessages{
				Type:       types.PostConsensusPartialSig,
				Slot:       cd.Duty.Slot,
				Signatures: []*types.PartialSignatureMessage{m},
			},
		}
	}
}

// SignBeaconObject signs a beacon object and returns the original objects slice appending partial sig message
func SignBeaconObject(runner *Runner, objects ...interface{}) (error, []interface{}) {
	//cd := objects[0].(*types.ConsensusData)
	// sign with domain
	return nil, append(objects, nil /* partial sig message */)
}

// QBFTProcessMessage process consensus message, returns:
// - Decided value if decided with quorum,
// - Stop if no quorum or previously decided
func QBFTProcessMessage(runner *Runner, objects ...interface{}) (error, []interface{}) {
	prevDecided := false
	if runner.qbft.Decided() {
		prevDecided = true
	}

	if err := runner.qbft.ProcessMessage(objects[0].(*qbft.SignedMessage)); err != nil {
		return err, nil
	}

	if !runner.qbft.Decided() || prevDecided {
		return nil, []interface{}{Stop}
	}

	return nil, []interface{}{runner.qbft.DecidedValue}
}

// DecodeMessage decodes a P2P message, error if can't
func DecodeMessage(runner *Runner, objects ...interface{}) (error, []interface{}) {
	msg := objects[0].(*p2p.Message)
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

// AddPostConsensusMessage adds post consensus msg to container
func AddPostConsensusMessage(runner *Runner, objects ...interface{}) (error, []interface{}) {
	return nil, objects
}

// ValidatePartialSignatureForSlot validates a provided post consensus message
func ValidatePartialSignatureForSlot(runner *Runner, objects ...interface{}) (error, []interface{}) {
	//slot :=  runner.qbft.DecidedValue.(*types.ConsensusData).Duty.Slot
	//
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot
	return nil, objects
}

// VerifyExpectedRoots validates a provided post consensus message
func VerifyExpectedRoots(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot with decided consensus
	return nil, objects
}
