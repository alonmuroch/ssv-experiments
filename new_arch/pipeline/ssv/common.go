package ssv

import (
	"ssv-experiments/new_arch/pipeline"
	qbft2 "ssv-experiments/new_arch/pipeline/qbft"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

const (
	PreConsensusPhase  = "PreConsensusPhase"
	ConsensusPhase     = "ConsensusPhase"
	PostConsensusPhase = "PostConsensusPhase"
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

// SignBeaconObject receives a consensus data and returns SignedPartialSignatureMessages
func SignBeaconObject(t types.PartialSigMsgType) func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
		cd := objects[0].(*types.ConsensusData)

		r, err := cd.GetSigningRoot()
		if err != nil {
			return err, nil
		}

		m := &types.SignedPartialSignatureMessages{
			Message: types.PartialSignatureMessages{
				Type:       t,
				Slot:       cd.Duty.Slot,
				Identifier: pipeline.Runner.Identifier,
				Signatures: []*types.PartialSignatureMessage{
					{
						Root:      r,
						Signature: [96]byte{},
					},
				},
			},
			Signer: 1,
		}
		// sign with domain
		return nil, []interface{}{m}
	}
}

// QBFTProcessMessage process consensus message, returns:
// - Decided ConsensusData if decided with quorum,
// - Stop if no quorum or previously decided
func QBFTProcessMessage(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	prevDecided := false
	if p.Instance.Decided() {
		prevDecided = true
	}

	qbftPipeline, err := qbft2.NewQBFTPipelineFromInstance(p.Instance)
	if err != nil {
		return err, nil
	}

	err, _ = qbftPipeline.ProcessMessage(objects[0].(*qbft.SignedMessage))
	if err != nil {
		return err, nil
	}

	if !p.Instance.Decided() || prevDecided {
		return nil, []interface{}{pipeline.Stop}
	}

	v, err := p.Instance.DecidedValue()
	if err != nil {
		return err, nil
	}

	ret := &types.ConsensusData{}
	if err := ret.UnmarshalSSZ(v); err != nil {
		return err, nil
	}

	return nil, []interface{}{ret}
}

// AddPostConsensusMessage adds post consensus msg to container
func AddPostConsensusMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	msg := objects[0].(*types.SignedPartialSignatureMessages)
	pipeline.Runner.State.PartialSignatures = append(pipeline.Runner.State.PartialSignatures, msg)
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

// NoQuorumStop checks if objects[0] == true, pass rest of ojects forward to next pipeline item. If false, stop
func NoQuorumStop(t types.PartialSigMsgType) func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
		if t.IsPostConsensusType() {
			if p.Runner.HasPostConsensusQuorum() {
				return nil, []interface{}{true}
			}
		} else {
			if p.Runner.HasPreConsensusQuorum() {
				return nil, []interface{}{true}
			}
		}
		return nil, []interface{}{pipeline.Stop}
	}
}

// NotQBFTMessageSkip checks if objects[0] is a qbft.SignedMessage, if not skip
func NotQBFTMessageSkip(nextPhase string) func(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
		// check if post consensus message
		_, ok := objects[0].(*qbft.SignedMessage)

		if ok { // consensus message
			return nil, objects
		}
		return nil, append(
			[]interface{}{
				pipeline.SkipToPhase,
				nextPhase,
			}, objects...)
	}
}

// NotPostConsensusMessageStop checks if objects[0] a post consensus partial sig message, if not stop
func NotPostConsensusMessageStop(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// check if post consensus message
	msg, ok := objects[0].(*types.SignedPartialSignatureMessages)

	if ok && msg.Message.Type.IsPostConsensusType() {
		return nil, objects
	}
	return nil, append(
		[]interface{}{
			pipeline.Stop,
		}, objects...)
}

// NotPreConsensusMessageSkip checks if objects[0] is a pre consensus partial sig message, if not skip
func NotPreConsensusMessageSkip(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// check if pre consensus message
	msg, ok := objects[0].(*types.SignedPartialSignatureMessages)

	if ok && msg.Message.Type.IsPreConsensusType() {
		return nil, objects
	}
	return nil, append(
		[]interface{}{
			pipeline.Stop,
		}, objects...)
}

// NotPreConsensusQuorumStop checks if pre consensus quorum, if not stop
func NotPreConsensusQuorumStop(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// TODO check pre-consensus quorum
	if false {
		return nil, []interface{}{pipeline.Stop}
	}
	return nil, objects
}

// NotDecidedStop checks if decided, if not stop
func NotDecidedStop(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	if !pipeline.Instance.Decided() {
		return nil, []interface{}{pipeline.Stop}
	}
	return nil, objects
}
