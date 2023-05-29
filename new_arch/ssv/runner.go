package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State      State
	Identifier p2p.Identifier
	pipeline   *Pipeline

	qbft *qbft.Instance
}

func NewRunner(duty *types.Duty) *Runner {
	return &Runner{
		State:      NewState(duty),
		Identifier: p2p.NewIdentifier(duty.Slot, duty.ValidatorPK, duty.Role),
		pipeline:   NewPipeline(),
	}
}

func (r *Runner) ProcessMessage(msg *p2p.Message) error {
	objsToPass := []interface{}{msg}
	err, _ := r.pipeline.Run(r, objsToPass)
	return err
}
