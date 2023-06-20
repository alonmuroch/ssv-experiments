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
	EndPhase         = "EndPhase"
)

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

// CreatePrepareMessage creates and returns prepare message encoded in p2p.Message
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
