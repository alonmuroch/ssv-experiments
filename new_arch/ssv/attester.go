package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

func NewAttesterRunnerForDuty(duty *types.Duty) *Runner {
	ret := StartAttesterRunner(NewRunner(duty))
	ret.pipeline.
		Add(DecodeMessage).
		AddIfConsensusMessage(
			// consensus phase
			NewPipeline().
				Add(QBFTProcessMessage).
				Add(ValidateDecidedValue(func(data *types.ConsensusData) error {
					return nil
				})).
				Add(SignBeaconObject).
				Add(ConstructPostConsensusMessage(types.PostConsensusPartialSig)).
				Add(Broadcast(p2p.SSVPartialSignatureMsgType)).
				Pipepify(),
		).
		AddIfPostConsensusMessage(
			// post-consensus phase
			NewPipeline().
				StopIfNotDecided().
				Add(ValidatePartialSignatureForSlot).
				Add(VerifyExpectedRoots).
				Add(AddPostConsensusMessage).
				Add(OnQuorumReconstructAttestationData).
				Add(BroadcastBeacon).
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

// ProcessAttesterConsensus will process attester consensus messages, wait for quorum and broadcast a partial signature over the decided value
func ProcessAttesterConsensus(runner *Runner, objects ...interface{}) (error, []interface{}) {
	err, objs := QBFTProcessMessage(runner, objects)
	if err != nil {
		return err, nil
	}

	if IsSkipNext(objs) {
		return nil, []interface{}{runner.qbft.DecidedValue}
	}

	if IsStop(objs) {
		return nil, objs
	}

	err, objs = ValidateDecidedValue(func(data *types.ConsensusData) error {
		return nil
	})(runner, objs)
	if err != nil {
		return err, nil
	}

	// sign attestation data
	// construct partial signature
	// broadcast
	return nil, []interface{}{}
}

// OnQuorumReconstructAttestationData checks quorum of post consensus msgs, reconstructs valid signed attestation and returns it. Otherwise stops
func OnQuorumReconstructAttestationData(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}
