package gas

import (
	"ssv-experiments/ssv_chain/operations"
)

const (
	ByteData        = 1
	OperatorAdd     = 10
	PublicKeyStore  = 5
	ModuleAdd       = 50
	DepositBalance  = 5
	WithdrawBalance = 10
)

// ConsumeGas consumes SSV gas for account, returns error if failed (insufficient, etc.)
func ConsumeGas(ctx *operations.Context, gas uint64) error {
	panic("implement")
}
