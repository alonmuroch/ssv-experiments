package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/ssv/pipeline"
	"ssv-experiments/new_arch/types"
)

func NewProposerRunnerForDuty(config Config, duty *types.Duty) *Runner {
	ret := StartProposerRunner(NewRunner(config, duty))
	ret.pipeline.
		Add(pipeline.DecodeMessage).

		// ##### pre consensus phase #####
		MarkPhase(pipeline.PreConsensusPhase).
		SkipIfNotPreConsensusMessage(pipeline.ConsensusPhase).
		Add(pipeline.ValidatePartialSignatureForSlot).
		Add(pipeline.VerifyExpectedRoots).
		Add(pipeline.AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(types.RandaoPartialSig).
		Add(ReconstructRandao).
		Add(FetchProposedBlock).
		Add(DecideOnBlock).

		// ##### pre consensus phase #####
		MarkPhase(pipeline.ConsensusPhase).
		SkipIfNotConsensusMessage(pipeline.PostConsensusPhase).
		StopINoPreConsensusQuorum().
		Add(pipeline.QBFTProcessMessage).
		Add(pipeline.ValidateDecidedValue(func(data *types.ConsensusData) error {
			return nil
		})).
		Add(pipeline.SignBeaconObject).
		Add(pipeline.ConstructPostConsensusMessage(types.PostConsensusPartialSig)).
		Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).

		// ##### post consensus phase #####
		MarkPhase(pipeline.PostConsensusPhase).
		StopIfNotPostConsensusMessage().
		StopIfNotDecided().
		Add(pipeline.ValidatePartialSignatureForSlot).
		Add(pipeline.VerifyExpectedRoots).
		Add(pipeline.AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(types.PostConsensusPartialSig).
		Add(ReconstructBlockData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)

	return ret
}

func StartProposerRunner(r *Runner) *Runner {
	// Get randao
	// broadcast partial sig for randao

	return r
}

// ReconstructBlockData reconstructs valid signed block and returns it.
func ReconstructBlockData(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// ReconstructRandao reconstructs returns it.
func ReconstructRandao(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// FetchProposedBlock takes as input reconstructed randao signature, add fetched proposed block
func FetchProposedBlock(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// DecideOnBlock takes as input proposed block data, constructs consensus data and starts a qbft instance
func DecideOnBlock(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no pre consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
