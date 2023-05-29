package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/ssv/pipeline"
	"ssv-experiments/new_arch/types"
)

func NewProposerRunnerForDuty(duty *types.Duty) *Runner {
	ret := StartProposerRunner(NewRunner(duty))
	ret.pipeline.
		Add(pipeline.DecodeMessage).
		AddIfPreConsensusMessage(
			pipeline.NewPipeline().
				Add(pipeline.ValidatePartialSignatureForSlot).
				Add(pipeline.VerifyExpectedRoots).
				Add(pipeline.AddPostConsensusMessage).
				Add(OnQuorumReconstructRandao).
				Add(FetchProposedBlock).
				Add(DecideOnBlock).
				Pipepify(),
		).
		StopINoPreConsensusQuorum().
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
				Add(OnQuorumReconstructBlockData).
				Add(pipeline.BroadcastBeacon).
				Pipepify(),
		)

	return ret
}

func StartProposerRunner(r *Runner) *Runner {
	// Get randao
	// broadcast partial sig for randao

	return r
}

// OnQuorumReconstructBlockData checks quorum of post consensus msgs, reconstructs valid signed block and returns it. Otherwise stops
func OnQuorumReconstructBlockData(runner *Runner, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// OnQuorumReconstructRandao checks quorum of randao partial msgs, reconstructs returns it. Otherwise stops
func OnQuorumReconstructRandao(runner *Runner, objects ...interface{}) (error, []interface{}) {
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
