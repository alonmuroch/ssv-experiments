package module

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type addModuleV0 struct {
	Name    []byte  `ssz-max:"1024"`
	Network [4]byte `ssz-size:"4"`
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	// OP_Add is a user event for creating a new module
	case types.OP_Add:
		opObj := &addModuleV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// consume gas
		if err := gas.ConsumeGas(ctx, gas.ModuleAdd+gas.ByteData*uint64(len(opObj.Name))); err != nil {
			return err
		}

		// Verify network exists, if not fail and consume gas
		if !ctx.Config.IsSupportedNetwork(opObj.Network) {
			return fmt.Errorf("network not found")
		}

		// update operators
		ctx.State.Modules = append(ctx.State.Modules, &types.Module{
			Address: ctx.Account.Address,
			ID:      uint64(len(ctx.State.Modules)),
			Name:    opObj.Name,
			Network: opObj.Network,
		})
		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}
