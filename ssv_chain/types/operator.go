package types

import "ssv-experiments/ssv_chain/common"

type Operator struct {
	// Account that controls the operator
	Account uint64
	// ID is unique for each operator
	ID uint64
	// PublicKey the operator uses to send messages
	PublicKey *common.CryptoKey
	// Modules IDs registered to
	Modules []uint64
}
