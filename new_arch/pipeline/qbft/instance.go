package qbft

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewQBFTPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	return NewPipeline(instance).
		// ##### propose if proposer #####
		Add(ProposeForFirstRound).

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

// ProposeForFirstRound will broadcast a proposal if first round and is proposer
func ProposeForFirstRound(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	if p.Instance.IsFirstRound() == p.Instance.IsProposer() {
		msg, err := p.Instance.CreateProposalMessage()
		if err != nil {
			return err, nil
		}
		err, _ = pipeline.Broadcast(p2p.SSVConsensusMsgType)(p, msg)
		if err != nil {
			return err, nil
		}
	}

	return nil, objects
}
