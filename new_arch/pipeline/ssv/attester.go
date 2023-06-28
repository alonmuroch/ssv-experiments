package ssv

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	qbft2 "ssv-experiments/new_arch/pipeline/qbft"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func NewAttesterPipeline(runner *ssv.Runner) (*pipeline.Pipeline, error) {
	ret := pipeline.NewPipeline()
	ret.Identifier = p2p.NewIdentifier(runner.State.StartingDuty.Slot, runner.State.StartingDuty.ValidatorPK, runner.State.StartingDuty.Role)
	ret.Runner = runner
	ret.
		// ##### fetch attestation data and start QBFT instance #####
		MarkPhase(pipeline.InitPhase).
		Add(DecideOnAttestationData).
		Stop().

		// ##### start #####
		MarkPhase(pipeline.StartPhase).
		Add(pipeline.ValidateP2PMessage).
		Add(pipeline.DecodeMessage).

		// ##### consensus phase #####
		MarkPhase(ConsensusPhase).
		Add(NotQBFTMessageSkip(PostConsensusPhase)).
		Add(QBFTProcessMessage).
		Add(ValidateDecidedValue(func(data *types.ConsensusData) error {
			return nil
		})).
		Add(SignBeaconObject(types.PostConsensusPartialSig)).
		Add(pipeline.Broadcast(p2p.SSVPartialSignatureMsgType)).
		Stop().

		// ##### post consensus phase #####
		MarkPhase(PostConsensusPhase).
		Add(NotPostConsensusMessageStop).
		Add(NotDecidedStop).
		Add(ValidatePartialSignatureForSlot).
		Add(VerifyExpectedRoots).
		Add(AddPostConsensusMessage).
		Add(NoQuorumStop(types.PostConsensusPartialSig)).
		Add(ReconstructAttestationData).
		Add(pipeline.BroadcastBeacon).

		// ##### end phase #####
		MarkPhase(pipeline.EndPhase)

	return ret, ret.Init()
}

// ReconstructAttestationData reconstructs valid signed attestation and returns it
func ReconstructAttestationData(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// if no post consensus quorum, return stop

	// iterate all roots, reconstruct signature and return
	return nil, []interface{}{}
}

// DecideOnAttestationData takes as input proposed attestation data, constructs consensus data and starts a qbft instance
func DecideOnAttestationData(pipeline *pipeline.Pipeline, objects ...interface{}) (error, []interface{}) {
	// already running
	if pipeline.Instance != nil {
		return nil, objects
	}

	// TODO - get attestation data
	attData := &phase0.AttestationData{}
	byts, err := attData.MarshalSSZ()
	if err != nil {
		return err, nil
	}

	inputData := &types.ConsensusData{
		Duty:        pipeline.Runner.State.StartingDuty,
		DataVersion: 0,
		DataSSZ:     byts,
	}

	pipeline.Instance = qbft.NewInstance(inputData, pipeline.Runner.Share, pipeline.Runner.State.StartingDuty.Slot, pipeline.Runner.State.StartingDuty.Role)
	// start the instance
	_, err = qbft2.NewQBFTPipelineFromInstance(
		pipeline.Instance,
		p2p.NewIdentifier(inputData.Duty.Slot, inputData.Duty.ValidatorPK, inputData.Duty.Role),
	)
	if err != nil {
		return err, nil
	}

	return nil, objects
}
