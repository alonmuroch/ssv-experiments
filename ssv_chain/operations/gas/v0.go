package gas

import (
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations"
)

const (
	ByteData       = common.VGBitOneHundredthSSV
	OperatorAdd    = common.VGBitOneSSV
	PublicKeyStore = common.VGBitTenthSSV * 5
	ModuleAdd      = common.VGBitOneSSV * 5

	// balance
	DepositBalance  = common.VGBitTenthSSV * 5
	WithdrawBalance = common.VGBitOneSSV

	// cluster
	ClusterAdd            = common.VGBitOneSSV * 5
	ClusterModify         = common.VGBitOneSSV * 5
	ClusterInstanceAdd    = common.VGBitTenthSSV * 5
	ClusterInstanceRemove = common.VGBitTenthSSV * 5

	// staking
	LockUnlockStake = common.VGBitTenthSSV * 5
	DelegateStake   = common.VGBitTenthSSV * 5

	// validator
	ValidatorAdd = common.VGBitOneSSV
)

// ConsumeGas consumes SSV gas for account, returns error if failed (insufficient, etc.)
func ConsumeGas(ctx *operations.Context, gas uint64) error {
	gasCost := ctx.GasCost(gas)
	ctx.GasConsumed += gas
	return ctx.Account.ReduceBalance(gasCost, ctx.Config.MainSSVTokenAddress, ctx.Config.MainSSVTokenNetwork)
}
