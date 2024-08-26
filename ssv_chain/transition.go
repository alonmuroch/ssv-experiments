package ssv_chain

import (
	"fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	"ssv-experiments/ssv_chain/operations"
	account "ssv-experiments/ssv_chain/operations/account"
	cluster "ssv-experiments/ssv_chain/operations/cluster"
	module "ssv-experiments/ssv_chain/operations/module"
	operator "ssv-experiments/ssv_chain/operations/operator"
	"ssv-experiments/ssv_chain/types"
)

// ProcessBlock is a CometBFT compatible request that is called during FinalizeBlock.
//
//	It applies the block to the state and returns a CometBFT compatible response
func ProcessBlock(state *types.State, req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	if err := ApplyBlockHeight(state, uint64(req.Height)); err != nil {
		return nil, err
	}
	panic("implement")
}

func ApplyBlockHeight(state *types.State, newHeight uint64) error {
	if state.BlockHeight+1 != newHeight {
		return fmt.Errorf("invalid height")
	}
	state.BlockHeight++
	return nil
}

type Receipt struct {
	GasConsumed uint64
	Error       error
}

// ProcessTransactions processes and applies transactions on state, returns nil if valid
func ProcessTransactions(s *types.State, txs []*types.Transaction) *Receipt {
	gasConsumed := uint64(0)
	for _, tx := range txs {
		ctx := &operations.Context{
			State:   s,
			Account: s.AccountByAddress(tx.Address),
		}
		if err := ProcessTransaction(ctx, tx); err != nil {
			return &Receipt{Error: err}
		}

		gasConsumed += ctx.GasConsumed
	}

	return &Receipt{GasConsumed: gasConsumed}
}

func ProcessTransaction(ctx *operations.Context, tx *types.Transaction) error {
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
	return nil
}
