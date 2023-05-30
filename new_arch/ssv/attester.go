package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/ssv/pipeline"
	"ssv-experiments/new_arch/types"
)

func NewAttesterRunnerForDuty(duty *types.Duty) *Runner {
	ret := StartAttesterRunner(NewRunner(duty))
	ret.pipeline.
		Add(pipeline.DecodeMessage).

		// ##### consensus phase #####
		MarkPhase(pipeline.ConsensusPhase).
		SkipIfNotConsensusMessage(pipeline.PostConsensusPhase).
		Add(pipeline.QBFTProcessMessage).
		Add(pipeline.ValidateDecidedValue(func(data *types.ConsensusData) error {
			return nil
		})).
		Add(pipeline.SignBeaconObject).
		Add(pipeline.ConstructPostConsensusMessage(types.PostConsensusPartialSig)).
		Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).

		// ##### post consensus phase #####
		MarkPhase(pipeline.PostConsensusPhase).
		SkipIfNotPostConsensusMessage(pipeline.EndPhase).
		StopIfNotDecided().
		Add(pipeline.ValidatePartialSignatureForSlot).
		Add(pipeline.VerifyExpectedRoots).
		Add(pipeline.AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(PostConsensus).
		Add(ReconstructAttestationData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)

	return ret
}

func StartAttesterRunner(r *Runner) *Runner {
	// Get attestation data
	// construct consensus data
	// start QBFT Instance

	return r
}

// ReconstructAttestationData reconstructs valid signed attestation and returns it
func ReconstructAttestationData(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
