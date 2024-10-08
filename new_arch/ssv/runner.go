package ssv

import (
	"ssv-experiments/new_arch/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State *State
	// share is the share for the runner with which messages are verified and signed
	Share *types.Share
}

func NewRunner(share *types.Share, duty *types.Duty) *Runner {
	return &Runner{
		State: NewState(duty),
		Share: share,
	}
}

func (r *Runner) HasPreConsensusQuorum() bool {
	all := r.State.PartialSignatures.AllPreConsensus()
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) HasPostConsensusQuorum() bool {
	all := r.State.PartialSignatures.AllPostConsensus()
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) Finished() bool {
	if r.State.DecidedConsensusData() == nil {
		return false
	}

	return r.HasPostConsensusQuorum()
}
