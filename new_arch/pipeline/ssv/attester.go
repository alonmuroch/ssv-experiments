package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func NewAttesterPipeline(runner *ssv.Runner) *pipeline.Pipeline {
	return pipeline.NewPipeline(runner).
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
		StopIfNotPostConsensusMessage().
		StopIfNotDecided().
		Add(pipeline.ValidatePartialSignatureForSlot).
		Add(pipeline.VerifyExpectedRoots).
		Add(pipeline.AddPostConsensusMessage).
		StopIfNoPartialSigQuorum(types.PostConsensusPartialSig).
		Add(ReconstructAttestationData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)
}

// ReconstructAttestationData reconstructs valid signed attestation and returns it
func ReconstructAttestationData(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
