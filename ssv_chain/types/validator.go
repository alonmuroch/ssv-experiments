package types

import (
	"fmt"
	"ssv-experiments/ssv_chain/common"
)

type Validator struct {
	// ID is unique for each validator
	ID uint64
	// Active is true when the validator is actively participating
	Active bool
	// Slashed when validator is slashed, exited from ever being an active validator
	Slashed bool
	// Address is an L1 address controlling the account
	Address []byte `ssz-max:"128"`
	// VotingPower represents the validator's voting power, per cometBFT.ValidatorUpdate struct
	VotingPower uint64
	// PublicKey the validator uses to vote
	PublicKey *common.CryptoKey
}

func (v *Validator) Penalize(state *State, config Configure, amount uint64) error {
	acc := state.AccountByAddress(v.Address)
	if acc == nil {
		return fmt.Errorf("account not found")
	}

	return acc.ReduceBalance(amount, config.MainSSVTokenAddress, config.MainSSVTokenNetwork)
}

func (v *Validator) Slash(state *State, config Configure) error {
	v.Active = false
	v.Slashed = true
	v.VotingPower = 0

	acc := state.AccountByAddress(v.Address)
	if acc == nil {
		return fmt.Errorf("account not found")
	}

	b := acc.BalanceByTokenAddress(config.MainSSVTokenAddress, config.MainSSVTokenNetwork)
	if b == nil {
		return fmt.Errorf("balance not found")
	}

	return acc.ReduceBalance(b.Amount, config.MainSSVTokenAddress, config.MainSSVTokenNetwork)
}
