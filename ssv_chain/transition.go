package ssv_chain

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	account "ssv-experiments/ssv_chain/operations/account"
	cluster "ssv-experiments/ssv_chain/operations/cluster"
	module "ssv-experiments/ssv_chain/operations/module"
	operator "ssv-experiments/ssv_chain/operations/operator"
	"ssv-experiments/ssv_chain/types"
)

type Receipt struct {
	GasConsumed uint64
	Error       error
}

// ProcessTransactions processes and applies transactions on state, returns nil if valid
// Transactions should be VALIDATED!! for signatures before being included
func ProcessTransactions(
	ctx *operations.Context,
	s *types.State,
	txs []*types.Transaction,
) *Receipt {
	gasConsumed := uint64(0)
	for _, tx := range txs {
		ctx.Account = s.AccountByAddress(tx.Address)
		if ctx.Account == nil {
			return &Receipt{Error: fmt.Errorf("account not found")}
		}

		if err := ProcessTransaction(ctx, tx); err != nil {
			return &Receipt{Error: err}
		}

		gasConsumed += ctx.GasConsumed
	}

	return &Receipt{GasConsumed: gasConsumed}
}

func ProcessTransaction(ctx *operations.Context, tx *types.Transaction) error {
	if ctx.Account.Nonce != tx.Nonce {
		return fmt.Errorf("wrong nonce")
	}

	for _, op := range tx.Operations {
		t := op.Type[1]
		v := op.Type[2]
		subOP := op.Type[3]
		switch t {
		case types.OP_Module:
			return module.ProcessOperation(ctx, v, subOP, op.OperationData)
		case types.OP_Cluster:
			return cluster.ProcessOperation(ctx, v, subOP, op.OperationData)
		case types.OP_Operator:
			return operator.ProcessOperation(ctx, v, subOP, op.OperationData)
		case types.OP_Account:
			return account.ProcessOperation(ctx, v, subOP, op.OperationData)
		default:
			return fmt.Errorf("unknown operation type: %v", t)
		}
	}

	ctx.Account.Nonce++

	return nil
}
