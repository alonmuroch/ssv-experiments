package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State *State
	// share is the share for the runner with which messages are verified and signed
	Share *types.Share
	// identifier identifies this particular runner
	Identifier p2p.Identifier `ssz-size:"56"`

	// qbft holds the qbft instance for this runner.
	// It is left outside the state as the state should change if and when decided (setting DecidedData), this is not strictly part of the runner's state
	qbft *qbft.Instance
}

func NewRunner(share *types.Share, duty *types.Duty) *Runner {
	return &Runner{
		State:      NewState(duty),
		Share:      share,
		Identifier: p2p.NewIdentifier(duty.Slot, share.ValidatorPubKey, duty.Role),
	}
}

func (r *Runner) GetQBFT() *qbft.Instance {
	return r.qbft
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
	if r.State.DecidedData == nil {
		return false
	}

	return r.HasPostConsensusQuorum()
}
