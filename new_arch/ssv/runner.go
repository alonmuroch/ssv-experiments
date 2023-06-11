package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	pipeline2 "ssv-experiments/new_arch/ssv/pipeline"
	"ssv-experiments/new_arch/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State State
	// pipeline is the entire message process pipeline for any runner messages
	pipeline *pipeline2.Pipeline
	config   Config

	// qbft holds the qbft instance for this runner.
	// It is left outside the state as the state should change if and when decided (setting DecidedData), this is not strictly part of the runner's state
	qbft *qbft.Instance
}

func NewRunner(config Config, duty *types.Duty) *Runner {
	return &Runner{
		State:    NewState(duty),
		config:   config,
		pipeline: pipeline2.NewPipeline(),
	}
}

func (r *Runner) GetQBFT() *qbft.Instance {
	return r.qbft
}

func (r *Runner) GetConfig() *Config {
	return &r.config
}

func (r *Runner) HasPreConsensusQuorum() bool {
	all := r.State.PartialSignatures.AllPreConsensus()
	return len(all) >= int(r.config.Share.Quorum)
}

func (r *Runner) HasPostConsensusQuorum() bool {
	all := r.State.PartialSignatures.AllPostConsensus()
	return len(all) >= int(r.config.Share.Quorum)
}

func (r *Runner) ProcessMessage(msg *p2p.Message) error {
	objsToPass := []interface{}{msg}
	err, _ := r.pipeline.Run(r, objsToPass)
	return err
}
