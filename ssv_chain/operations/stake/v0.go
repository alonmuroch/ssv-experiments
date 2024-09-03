package stake

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type lockStakeV0 struct {
	Amounts []*types.Balance `ssz-max:"12"`
}

type validatorDelegationV0 struct {
	Amount      uint64 // positive to delegate and negative to un-delegate
	Sign        bool   // true = positive, false = negative
	ValidatorID uint64
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	case types.OP_Lock, types.OP_Release:
		opObj := &lockStakeV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		for _, a := range opObj.Amounts {
			// calculate and consume gas
			estimatedGas := uint64(gas.LockUnlockStake)
			estimatedGasCost := ctx.GasCost(estimatedGas)
			if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
				return err
			}
			ctx.GasConsumed += estimatedGas

			if !ctx.Config.ValidSSVTokenAddress(a.Network, a.TokenAddress) {
				continue
			}

			if op == types.OP_Lock {
				if err := ctx.Account.LockBalance(a.Amount, a.TokenAddress, a.Network); err != nil {
					return err
				}
			} else { // release
				if err := ctx.Account.ReleaseBalance(a.Amount, a.TokenAddress, a.Network); err != nil {
					return err
				}
			}
		}
		return nil
	case types.OP_Delegate:
		opObj := &validatorDelegationV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.DelegateStake)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// verify validator ID
		val := ctx.State.ValidatorByID(opObj.ValidatorID)
		if val == nil {
			return fmt.Errorf("validator not found")
		}

		// verify can delegate
		b := ctx.Account.BalanceByTokenAddress(ctx.Config.MainSSVTokenAddress, ctx.Config.MainSSVTokenNetwork)
		if b == nil {
			return fmt.Errorf("no SSV balance")
		}

		// validate amount
		/*
			    invalid | un-delegate	| delegate	|invalid
				////////|				|			|/////
				/_______|_______________|___________|____/
				////////|				|			|/////
						0			delegated  	  locked
		*/
		amount := int64(opObj.Amount)
		if !opObj.Sign {
			amount = int64(opObj.Amount) * -1
		}
		postB := int64(b.Delegated) + amount
		sufficientLocked := postB <= int64(b.Locked)
		sufficientDelegated := postB >= 0
		if !sufficientLocked {
			return fmt.Errorf("insufficient locked balance")
		}
		if !sufficientDelegated {
			return fmt.Errorf("insufficient delegated balance")
		}

		// update
		b.Delegated = uint64(postB)
		val.VotingPower = uint64(int64(val.VotingPower) + amount)

		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}
