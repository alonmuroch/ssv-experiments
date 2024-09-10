package account

import (
	"bytes"
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

type withdrawBalanceV0 struct {
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

		// add balances
		for _, b := range opObj.Balances {
			// consume gas
			if err := gas.ConsumeGas(ctx, gas.DepositBalance); err != nil {
				// roll back adding balance
				ctx.Account.ReduceBalance(b.Amount, b.TokenAddress, b.Network)
				return err
			}

			if !bytes.Equal(b.Network[:], ctx.Account.Network[:]) {
				return fmt.Errorf("not account network")
			}
			ctx.Account.AddBalance(b.Amount, b.TokenAddress, b.Network)
		}

		return nil
	case types.OP_Withdraw:
		opObj := &withdrawBalanceV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// consume gas
		if err := gas.ConsumeGas(ctx, gas.WithdrawBalance*uint64(len(opObj.Balances))); err != nil {
			return err
		}

		for _, b := range opObj.Balances {
			// reduce balance
			if err := ctx.Account.ReduceBalance(b.Amount, b.TokenAddress, b.Network); err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}
