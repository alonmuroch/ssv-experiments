package account

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type depositBalanceV0 struct {
	// Address is an L1 address controlling the account
	Address  []byte           `ssz-max:"128"`
	Balances []*types.Balance `ssz-max:"12"`
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	// OP_Deposit is a system operation that picks up deposit events on L1 and includes them as operations.
	// implementations should VERIFY!! every deposit operation matches an actual finalized L1 deposit event
	case types.OP_Deposit:
		opObj := &depositBalanceV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// attach account to context
		ctx.Account = ctx.State.AccountByAddress(opObj.Address)
		if ctx.Account == nil {
			ctx.Account = ctx.State.CreateAccountForAddress(opObj.Address)
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.DepositBalance)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		// add balances
		for _, b := range opObj.Balances {
			ctx.Account.DepositBalance(b)

			// consume gas
			if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
				// roll back adding balance
				ctx.Account.ReduceBalance(b.Amount, b.TokenAddress, b.Network)
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}
