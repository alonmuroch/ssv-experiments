package qbft

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewQBFTPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	return NewPipeline(instance).
		Add(pipeline.DecodeMessage).
		SkipIfNotConsensusMessage(EndPhase).

		// ##### proposal phase #####
		MarkPhase(ProposalPhase).
		SkipIfNotQBFTMessageType(PreparePhase, qbft.ProposalMessageType).
		Add(UponProposal).
		Add(CreatePrepareMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).

		// ##### prepare phase #####
		MarkPhase(PreparePhase).
		SkipIfNotQBFTMessageType(CommitPhase, qbft.PrepareMessageType).
		Add(UponPrepare).
		Add(NoQuorumStop).
		Add(CreateCommitMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).

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
