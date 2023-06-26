package qbft

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

const (
	ProposalPhase    = "ProposalPhase"
	PreparePhase     = "PreparePhase"
	CommitPhase      = "CommitPhase"
	RoundChangePhase = "RoundChangePhase"
)

// ValidateConsensusMessage validates consensus message (type, struct, etc), returns error if not valid
func ValidateConsensusMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	_, ok := objects[0].(*qbft.SignedMessage)
	if ok {
		return nil, objects
	}
	return errors.New("not a consensus message"), nil
}

// SignMessage receives an unsigned qbft.Message and returns a signed message
func SignMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	msg, ok := objects[0].(*qbft.Message)
	if !ok {
		return errors.New("object type not qbft.Message"), nil
	}

	ret := &qbft.SignedMessage{
		Message:   *msg,
		Signers:   []uint64{1},
		Signature: [96]byte{},
	}

	return nil, []interface{}{ret}
}

// UponPrepare runs the qbft.UponPrepare stage. Returns true and original objects slice if quorum reached
func UponPrepare(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	if err := pipeline.Instance.UponPrepare(objects[0].(*qbft.SignedMessage)); err != nil {
		return err, nil
	}
	return nil, append(
		[]interface{}{pipeline.Instance.PrepareQuorum()},
		objects...,
	)
}

// CreatePrepareMessage creates and returns prepare message
func CreatePrepareMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	msg, err := pipeline.Instance.CreatePrepareMessage()
	if err != nil {
		return err, nil
	}
	return nil, []interface{}{msg}
}

// UponProposal runs the qbft.UponProposal stage
func UponProposal(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	if err := pipeline.Instance.UponProposal(objects[0].(*qbft.SignedMessage)); err != nil {
		return err, nil
	}

	return nil, objects
}

// UponCommit runs the qbft.UponCommit stage. Returns true and original objects slice if quorum reached
func UponCommit(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	if err := pipeline.Instance.UponCommit(objects[0].(*qbft.SignedMessage)); err != nil {
		return err, nil
	}

	return nil, append(
		[]interface{}{pipeline.Instance.CommitQuorum()},
		objects...,
	)
}

// CreateCommitMessage creates and returns commit message
func CreateCommitMessage(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	msg, err := pipeline.Instance.CreateCommitMessage()
	if err != nil {
		return err, nil
	}
	return nil, []interface{}{msg}
}

// NoQuorumStop checks if objects[0] == true, pass rest of ojects forward to next pipeline item. If false, stop
func NoQuorumStop(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	quorum, ok := objects[0].(bool)
	if !ok {
		return errors.New("invalid input"), nil
	}

	if !quorum {
		return nil, []interface{}{pipeline.Stop}
	}
	return nil, objects[1:]
}

// NotQBFTMessageTypeSkip will validate message type, will skip if not
func NotQBFTMessageTypeSkip(nextPhase string, msgType uint64) func(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	return func(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
		msg, ok := objects[0].(*qbft.SignedMessage)
		if !ok {
			return nil, append(
				[]interface{}{
					pipeline.SkipToPhase,
					nextPhase,
				}, objects...)
		}

		if msg.Message.MsgType == msgType {
			return nil, objects
		}
		return nil, append(
			[]interface{}{
				pipeline.SkipToPhase,
				nextPhase,
			}, objects...)
	}
}
