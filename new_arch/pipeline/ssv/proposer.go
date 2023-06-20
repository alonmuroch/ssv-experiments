package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func NewProposerRunnerForDuty(runner *ssv.Runner) *pipeline.Pipeline {
	return NewPipeline(runner).
		Add(DecodeMessage).

		// ##### pre consensus phase #####
		MarkPhase(pipeline.PreConsensusPhase).
		SkipIfNotPreConsensusMessage(pipeline.ConsensusPhase).
		Add(ValidatePartialSignatureForSlot).
		Add(VerifyExpectedRoots).
		Add(AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(types.RandaoPartialSig).
		Add(ReconstructRandao).
		Add(FetchProposedBlock).
		Add(DecideOnBlock).

		// ##### consensus phase #####
		MarkPhase(pipeline.ConsensusPhase).
		SkipIfNotConsensusMessage(pipeline.PostConsensusPhase).
		StopINoPreConsensusQuorum().
		Add(QBFTProcessMessage).
		Add(ValidateDecidedValue(func(data *types.ConsensusData) error {
			return nil
		})).
		Add(SignBeaconObject).
		Add(ConstructPostConsensusMessage(types.PostConsensusPartialSig)).
		Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).

		// ##### post consensus phase #####
		MarkPhase(pipeline.PostConsensusPhase).
		StopIfNotPostConsensusMessage().
		StopIfNotDecided().
		Add(ValidatePartialSignatureForSlot).
		Add(VerifyExpectedRoots).
		Add(AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(types.PostConsensusPartialSig).
		Add(ReconstructBlockData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)
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
