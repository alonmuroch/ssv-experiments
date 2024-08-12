package account

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/types"
)

func ProcessAccountOperation(ctx *operations.Context, op, v byte, raw []byte) error {
	switch v {
	case types.OP_V0:
		return processAccountOperation(ctx, op, raw)
	default:
		return fmt.Errorf("unknown version %d", v)
	}
}
