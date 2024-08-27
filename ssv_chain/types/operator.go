package types

import "ssv-experiments/ssv_chain/common"

type PriceTier struct {
	Network  [4]byte `ssz-size:"4"`
	Capacity uint16
	// PricePerToken is how many payable tokens will be paid, per block, for each unit of network token (e.g. ethereum)
	PricePerToken uint64
	// PayableTokenAddress is the L1 address of the token paid for this tier
	PayableTokenAddress []byte `ssz-max:"64"`
	// WhitelistedAddress that can register to this tier, if empty any address can
	WhitelistedAddress [][]byte `ssz-max:"64,128"`
}

type Operator struct {
	// Address that controls the operator
	Address []byte `ssz-max:"128"`
	// ID is unique for each operator
	ID uint64
	// PublicKey the operator uses to send messages
	PublicKey *common.CryptoKey
	// Modules IDs registered to
	Module uint64
	// Tiers represent pricing tiers
	Tiers []*PriceTier `ssz-max:"16"`
}
