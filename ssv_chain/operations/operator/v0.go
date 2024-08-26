package operator

import (
	"fmt"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type addOperatorV0 struct {
	PublicKey *common.CryptoKey
	Module    uint64
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	case types.OP_Add:
		opObj := &addOperatorV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		gas := uint64(gas.OperatorAdd + gas.PublicKeyStore)
		estimatedGasCost := ctx.GasCost(gas)
		if ctx.Account.Balance < estimatedGasCost {
			return fmt.Errorf("insufficient gas")
		}

		// update gas
		ctx.Account.Balance -= estimatedGasCost
		ctx.GasConsumed = gas

		// Verify module exists, if not fail and consume gas
		if ctx.State.ModuleByID(opObj.Module) == nil {
			return fmt.Errorf("module not found")
		}

		// update operators
		ctx.State.Operators = append(ctx.State.Operators, &types.Operator{
			Account:   ctx.Account.ID,
			ID:        uint64(len(ctx.State.Operators)),
			PublicKey: opObj.PublicKey,
			Module:    opObj.Module,
		})
		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}
