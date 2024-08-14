package operations

import "ssv-experiments/ssv_chain/types"

type Context struct {
	// State is the pre-state before executing the tx
	State *types.State
	// Account executing the tx
	Account *types.Account
	// GasPrice set by the user when broadcasting the transaction
	GasPrice uint64
	// GasConsumed is set when tx applied to state
	GasConsumed uint64

	Error error
}

func (ctx *Context) GasCost(gas uint64) uint64 {
	return gas * ctx.GasPrice
}
