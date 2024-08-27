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
	Tiers     []*types.PriceTier `ssz-max:"16"`
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	case types.OP_Add:
		opObj := &addOperatorV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.OperatorAdd + gas.PublicKeyStore)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// Verify module exists, if not fail and consume gas
		if ctx.State.ModuleByID(opObj.Module) == nil {
			return fmt.Errorf("module not found")
		}

		// Validate price tiers, if not fail and consume gas
		if err := validateV0PriceTiers(opObj.Tiers); err != nil {
			return err
		}

		// update operators
		ctx.State.Operators = append(ctx.State.Operators, &types.Operator{
			Address:   ctx.Account.Address,
			ID:        uint64(len(ctx.State.Operators)),
			PublicKey: opObj.PublicKey,
			Module:    opObj.Module,
			Tiers:     opObj.Tiers,
		})
		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}

func validateV0PriceTiers(tiers []*types.PriceTier) error {
	if len(tiers) == 0 {
		return fmt.Errorf("no price tiers")
	}

	for _, t := range tiers {
		// Verify network exists, if not fail and consume gas
		if !common.IsSupportedNetwork(t.Network) {
			return fmt.Errorf("network not found")
		}
	}

	return nil
}
