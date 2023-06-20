package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func NewAttesterPipeline(runner *ssv.Runner) *pipeline.Pipeline {
	return NewPipeline(runner).
		Add(DecodeMessage).

		// ##### consensus phase #####
		MarkPhase(pipeline.ConsensusPhase).
		SkipIfNotConsensusMessage(pipeline.PostConsensusPhase).
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
