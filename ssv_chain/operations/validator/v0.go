package validator

import (
	"fmt"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type addValidatorV0 struct {
	PublicKey *common.CryptoKey
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	case types.OP_Add:
		opObj := &addValidatorV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.ValidatorAdd + gas.PublicKeyStore)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// update operators
		ctx.State.Validators = append(ctx.State.Validators, &types.Validator{
			Address:   ctx.Account.Address,
			ID:        uint64(len(ctx.State.Operators)),
			PublicKey: opObj.PublicKey,
			Active:    true,
		})
		return nil

	default:
		return fmt.Errorf("unknown operation")
	}
}
