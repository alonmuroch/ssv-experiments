package pipeline

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

// ValidateDecidedValue returns a pipeline function for a specific value check function
func ValidateDecidedValue(valueCheck func(data *types.ConsensusData) error) PipelineF {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		if err := valueCheck(objects[0].(*types.ConsensusData)); err != nil {
			return err, nil
		}
		return nil, objects
	}
}

// ConstructPostConsensusMessage receives consensus data and partial sig message and returns PartialSignatureMessages
func ConstructPostConsensusMessage(t types.PartialSigMsgType) func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
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
func SignBeaconObject(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	//cd := objects[0].(*types.ConsensusData)
	// sign with domain
	return nil, append(objects, nil /* partial sig message */)
}

// QBFTProcessMessage process consensus message, returns:
// - Decided value if decided with quorum,
// - Stop if no quorum or previously decided
func QBFTProcessMessage(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	prevDecided := false
	if runner.GetQBFT().Decided() {
		prevDecided = true
	}

	msgToBroadcast, err := runner.GetQBFT().ProcessMessage(objects[0].(*qbft.SignedMessage))
	if err != nil {
		return err, nil
	}

	if msgToBroadcast != nil {
		err, _ := Broadcast(p2p.SSVConsensusMsgType)(runner, msgToBroadcast)
		if err != nil {
			return err, nil
		}
	}

	if !runner.GetQBFT().Decided() || prevDecided {
		return nil, []interface{}{Stop}
	}

	return nil, []interface{}{runner.GetQBFT().DecidedValue}
}

// DecodeMessage decodes a P2P message, error if can't
func DecodeMessage(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
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
func AddPostConsensusMessage(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	return nil, objects
}

// ValidatePartialSignatureForSlot validates a provided post consensus message
func ValidatePartialSignatureForSlot(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	//slot :=  runner.qbft.DecidedValue.(*types.ConsensusData).Duty.Slot
	//
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot
	return nil, objects
}

// VerifyExpectedRoots validates a provided post consensus message
func VerifyExpectedRoots(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot with decided consensus
	return nil, objects
}
