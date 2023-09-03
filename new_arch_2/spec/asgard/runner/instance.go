package runner

import (
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State *types.State
	// share is the share for the runner with which messages are verified and signed
	Share *types.Share
}

func ProcessMessage(state *types.State, message *types.SignedPartialSignatureMessages) error {
	panic("implement")
}

func (r *Runner) HasPreConsensusQuorum() bool {
	all := AllPreConsensus(r.State)
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) HasPostConsensusQuorum() bool {
	all := AllPostConsensus(r.State)
	return len(all) >= int(r.Share.Quorum)
}

func (r *Runner) Finished() bool {
	if DecidedConsensusData(r.State) == nil {
		return false
	}

	return r.HasPostConsensusQuorum()
}
