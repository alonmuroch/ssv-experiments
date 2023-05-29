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
		AddIfConsensusMessage(
			// consensus phase
			pipeline.NewPipeline().
				Add(pipeline.QBFTProcessMessage).
				Add(pipeline.ValidateDecidedValue(func(data *types.ConsensusData) error {
					return nil
				})).
				Add(pipeline.SignBeaconObject).
				Add(pipeline.ConstructPostConsensusMessage(types.PostConsensusPartialSig)).
				Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).
				Pipepify(),
		).
		StopIfNotDecided().
		AddIfPostConsensusMessage(
			// post-consensus phase
			pipeline.NewPipeline().
				Add(pipeline.ValidatePartialSignatureForSlot).
				Add(pipeline.VerifyExpectedRoots).
				Add(pipeline.AddPostConsensusMessage).
				Add(OnQuorumReconstructAttestationData).
				Add(pipeline.BroadcastBeacon).
				Pipepify(),
		)

	return ret
}

func StartAttesterRunner(r *Runner) *Runner {
	// Get attestation data
	// construct consensus data
	// start QBFT Instance

	return r
}

// OnQuorumReconstructAttestationData checks quorum of post consensus msgs, reconstructs valid signed attestation and returns it. Otherwise stops
func OnQuorumReconstructAttestationData(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
