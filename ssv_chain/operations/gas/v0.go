package gas

import (
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations"
)

const (
	ByteData              = common.OneHundredthSSV
	OperatorAdd           = common.OneSSV
	PublicKeyStore        = common.TenthSSV * 5
	ModuleAdd             = common.OneSSV * 5
	DepositBalance        = common.TenthSSV * 5
	WithdrawBalance       = common.OneSSV
	ClusterAdd            = common.OneSSV * 5
	ClusterModify         = common.OneSSV * 5
	ClusterInstanceAdd    = common.TenthSSV * 5
	ClusterInstanceRemove = common.TenthSSV * 5
	LockUnlockStake       = common.TenthSSV * 5
)

// ConsumeGas consumes SSV gas for account, returns error if failed (insufficient, etc.)
func ConsumeGas(ctx *operations.Context, gas uint64) error {
	panic("implement")
}
