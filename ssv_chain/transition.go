package ssv_chain

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	account "ssv-experiments/ssv_chain/operations/account"
	cluster "ssv-experiments/ssv_chain/operations/cluster"
	module "ssv-experiments/ssv_chain/operations/module"
	operator "ssv-experiments/ssv_chain/operations/operator"
	"ssv-experiments/ssv_chain/operations/stake"
	"ssv-experiments/ssv_chain/operations/validator"
	"ssv-experiments/ssv_chain/types"
)

type Receipt struct {
	GasPerTx    []uint64
	GasConsumed uint64
	Errors      []error
}

// ProcessTransactions processes and applies transactions on state, returns nil if valid
// Transactions should be !!VALIDATED!! for signatures before being included
func ProcessTransactions(
	ctx *operations.Context,
	s *types.State,
	txs []*types.Transaction,
) *Receipt {
	gasConsumed := uint64(0)
	gasPerTx := make([]uint64, len(txs))
	errors := make([]error, len(txs))
	for i, tx := range txs {
		ctx.Account = s.AccountByAddress(tx.Address)
		if ctx.Account == nil {
			errors[i] = fmt.Errorf("account not found")
			continue
		}

		if err := ProcessTransaction(ctx, tx); err != nil {
			errors[i] = err
		}

		gasPerTx[i] = ctx.GasConsumed
		gasConsumed += ctx.GasConsumed
	}

	return &Receipt{
		GasPerTx:    gasPerTx,
		GasConsumed: gasConsumed,
		Errors:      errors,
	}
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
		case types.OP_Stake:
			return stake.ProcessOperation(ctx, v, subOP, op.OperationData)
		case types.OP_Validator:
			return validator.ProcessOperation(ctx, v, subOP, op.OperationData)
		default:
			return fmt.Errorf("unknown operation type: %v", t)
		}
	}

	ctx.Account.Nonce++

	return nil
}
