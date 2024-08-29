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
	default:
		return fmt.Errorf("unknown operation")
	}
}
