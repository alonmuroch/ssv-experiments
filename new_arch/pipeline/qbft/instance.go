package qbft

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

func NewQBFTPipeline(inputData *types.ConsensusData, share *types.Share) (*pipeline.Pipeline, error) {
	instance := qbft.NewInstance(inputData, share, inputData.Duty.Role, inputData.Duty.Role)
	return NewQBFTPipelineFromInstance(instance, p2p.NewIdentifier(inputData.Duty.Slot, inputData.Duty.ValidatorPK, inputData.Duty.Role))
}

func NewQBFTPipelineFromInstance(instance *qbft.Instance, identifier p2p.Identifier) (*pipeline.Pipeline, error) {
	ret := pipeline.NewPipeline()
	ret.Identifier = identifier
	ret.Instance = instance
	ret.
		// ##### propose if proposer #####
		MarkPhase(pipeline.InitPhase).
		Add(ProposeForFirstRound).
		Stop().

		// ##### start #####
		MarkPhase(pipeline.StartPhase).
		Add(ValidateConsensusMessage).

		// ##### proposal phase #####
		MarkPhase(ProposalPhase).
		Add(NotQBFTMessageTypeSkip(PreparePhase, qbft.ProposalMessageType)).
		Add(UponProposal).
		Add(CreatePrepareMessage).
		Add(SignMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).
		Stop().

		// ##### prepare phase #####
		MarkPhase(PreparePhase).
		Add(NotQBFTMessageTypeSkip(CommitPhase, qbft.PrepareMessageType)).
		Add(UponPrepare).
		Add(NoQuorumStop).
		Add(CreateCommitMessage).
		Add(SignMessage).
		Add(pipeline.Broadcast(p2p.SSVConsensusMsgType)).
		Stop().

		// ##### commit phase #####
		MarkPhase(CommitPhase).
		Add(NotQBFTMessageTypeSkip(RoundChangePhase, qbft.CommitMessageType)).
		Add(UponCommit).
		Add(NoQuorumStop).
		Stop().

		// ##### round change phase #####
		MarkPhase(RoundChangePhase).
		Add(NotQBFTMessageTypeSkip(pipeline.EndPhase, qbft.RoundChangeMessageType)).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase).
		Stop()

	return ret, ret.Init() // runs init phase
}

// ProposeForFirstRound will broadcast a proposal if first round and is proposer
func ProposeForFirstRound(p *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// TODO - broadcast only once
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
