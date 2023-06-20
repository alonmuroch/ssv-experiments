package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	qbft2 "ssv-experiments/new_arch/pipeline/qbft"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

const (
	PreConsensusPhase  = "PreConsensusPhase"
	ConsensusPhase     = "ConsensusPhase"
	PostConsensusPhase = "PostConsensusPhase"
	EndPhase           = "EndPhase"
)

// ValidateDecidedValue returns a pipeline function for a specific value check function
func ValidateDecidedValue(valueCheck func(data *types.ConsensusData) error) pipeline.PipelineF {
	return func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
		if err := valueCheck(objects[0].(*types.ConsensusData)); err != nil {
			return err, nil
		}
		return nil, objects
	}
}

// ConstructPostConsensusMessage receives consensus data and partial sig message and returns PartialSignatureMessages
func ConstructPostConsensusMessage(t types.PartialSigMsgType) func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
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
func SignBeaconObject(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	//cd := objects[0].(*types.ConsensusData)
	// sign with domain
	return nil, append(objects, nil /* partial sig message */)
}

// QBFTProcessMessage process consensus message, returns:
// - Decided value if decided with quorum,
// - Stop if no quorum or previously decided
func QBFTProcessMessage(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	prevDecided := false
	if p.Instance.Decided() {
		prevDecided = true
	}

	qbftPipeline := qbft2.NewQBFTPipeline(p.Instance)
	err, msgToBroadcast := qbftPipeline.ProcessMessage(objects[0].(*qbft.SignedMessage))
	if err != nil {
		return err, nil
	}

	if msgToBroadcast != nil {
		err, _ := pipeline.Broadcast(p2p.SSVConsensusMsgType)(p, msgToBroadcast)
		if err != nil {
			return err, nil
		}
	}

	if !p.Instance.Decided() || prevDecided {
		return nil, []interface{}{pipeline.Stop}
	}

	return nil, []interface{}{p.Instance.DecidedValue}
}

// AddPostConsensusMessage adds post consensus msg to container
func AddPostConsensusMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return nil, objects
}

// ValidatePartialSignatureForSlot validates a provided post consensus message
func ValidatePartialSignatureForSlot(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	//slot :=  runner.qbft.DecidedValue.(*types.ConsensusData).Duty.Slot
	//
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot
	return nil, objects
}

// VerifyExpectedRoots validates a provided post consensus message
func VerifyExpectedRoots(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// verify objects[0].(*types.SignedPartialSignatureMessages) with slot with decided consensus
	return nil, objects
}
