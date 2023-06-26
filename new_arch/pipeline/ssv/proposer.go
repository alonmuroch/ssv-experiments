package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func NewProposerRunnerForDuty(runner *ssv.Runner) (*pipeline.Pipeline, error) {
	ret := pipeline.NewPipeline()
	ret.Runner = runner
	ret.
		Add(pipeline.DecodeMessage).

		// ##### pre consensus phase #####
		MarkPhase(PreConsensusPhase).
		Add(NotPreConsensusMessageSkip).
		Add(ValidatePartialSignatureForSlot).
		Add(VerifyExpectedRoots).
		Add(AddPostConsensusMessage).
		Add(NoQuorumStop(types.RandaoPartialSig)).
		Add(ReconstructRandao).
		Add(FetchProposedBlock).
		Add(DecideOnBlock).

		// ##### consensus phase #####
		MarkPhase(ConsensusPhase).
		Add(NotQBFTMessageSkip(PostConsensusPhase)).
		Add(NotPreConsensusQuorumStop).
		Add(QBFTProcessMessage).
		Add(ValidateDecidedValue(func(data *types.ConsensusData) error {
			return nil
		})).
		Add(SignBeaconObject(types.PostConsensusPartialSig)).
		Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).

		// ##### post consensus phase #####
		MarkPhase(PostConsensusPhase).
		Add(NotPostConsensusMessageStop).
		Add(NotDecidedStop).
		Add(ValidatePartialSignatureForSlot).
		Add(VerifyExpectedRoots).
		Add(AddPostConsensusMessage).
		Add(NoQuorumStop(types.PostConsensusPartialSig)).
		Add(ReconstructBlockData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)
	return ret, ret.Init()
}

// ReconstructBlockData reconstructs valid signed block and returns it.
func ReconstructBlockData(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// ReconstructRandao reconstructs returns it.
func ReconstructRandao(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// FetchProposedBlock takes as input reconstructed randao signature, add fetched proposed block
func FetchProposedBlock(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// DecideOnBlock takes as input proposed block data, constructs consensus data and starts a qbft instance
func DecideOnBlock(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
