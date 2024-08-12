package ssv_chain

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	module3 "ssv-experiments/ssv_chain/operations/account"
	module2 "ssv-experiments/ssv_chain/operations/cluster"
	"ssv-experiments/ssv_chain/operations/module"
	module4 "ssv-experiments/ssv_chain/operations/operator"
	"ssv-experiments/ssv_chain/types"
)

// ProcessTransactions processes and applies transactions on state, returns nil if valid
func ProcessTransactions(s *types.State, txs []*types.Transaction) error {
	for _, tx := range txs {
		t := tx.Type[1]
		v := tx.Type[2]
		op := tx.Type[3]
		ctx := &operations.Context{
			State:   s,
			Account: s.AccountByAddress(tx.Address),
		}
		switch t {
		case types.OP_Module:
			return module.ProcessModuleOperation(ctx, v, op, tx.OperationData)
		case types.OP_Cluster:
			return module2.ProcessClusterOperation(ctx, v, op, tx.OperationData)
		case types.OP_Operator:
			return module4.ProcessOperatorOperation(ctx, v, op, tx.OperationData)
		case types.OP_Account:
			return module3.ProcessAccountOperation(ctx, v, op, tx.OperationData)
		default:
			return fmt.Errorf("unknown transaction type: %v", t)
		}
	}

	return nil
}
