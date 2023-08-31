package runner

import (
	types2 "ssv-experiments/new_arch_2/spec/asgard/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State *State
	// share is the share for the runner with which messages are verified and signed
	Share *types2.Share
}

func NewRunner(share *types2.Share, duty *types2.Duty) *Runner {
	return &Runner{
		State: NewState(duty),
		Share: share,
	}
}

func (r *Runner) HasPreConsensusQuorum() bool {
	all := r.State.AllPreConsensus()
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) HasPostConsensusQuorum() bool {
	all := r.State.AllPostConsensus()
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) Finished() bool {
	if r.State.DecidedConsensusData() == nil {
		return false
	}

	return r.HasPostConsensusQuorum()
}
