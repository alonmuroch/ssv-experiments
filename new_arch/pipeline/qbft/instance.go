package qbft

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewQBFTPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	return NewPipeline(instance).
		// ##### start qbft instance #####
		DoOnce("start_qbft", StartQBFT, SignMessage, pipeline.Broadcast(p2p.SSVConsensusMsgType)).

		// ##### start #####
		Add(pipeline.DecodeMessage).
		SkipIfNotConsensusMessage(EndPhase).

		// ##### proposal phase #####
		MarkPhase(ProposalPhase).
		SkipIfNotQBFTMessageType(PreparePhase, qbft.ProposalMessageType).
		Add(UponProposal).
		Add(CreatePrepareMessage).
		Add(SignMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).
		Stop().

		// ##### prepare phase #####
		MarkPhase(PreparePhase).
		SkipIfNotQBFTMessageType(CommitPhase, qbft.PrepareMessageType).
		Add(UponPrepare).
		Add(NoQuorumStop).
		Add(CreateCommitMessage).
		Add(SignMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).
		Stop().

		// ##### commit phase #####
		MarkPhase(CommitPhase).
		SkipIfNotQBFTMessageType(RoundChangePhase, qbft.CommitMessageType).
		Add(UponCommit).
		Add(NoQuorumStop).
		Stop().

		// ##### round change phase #####
		MarkPhase(RoundChangePhase).
		SkipIfNotQBFTMessageType(EndPhase, qbft.RoundChangeMessageType).

		// ##### end phase #####
		MarkPhase(EndPhase)
}

// StartQBFT calls qbft.Instance.Start, returns message to broadcast or stops
func StartQBFT(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	msg, err := p.Instance.Start()
	if err != nil {
		if err != nil {
			return err, nil
		}
	}
	if msg == nil {
		return nil, []interface{}{pipeline.Stop}
	}
	return nil, []interface{}{msg}
}
