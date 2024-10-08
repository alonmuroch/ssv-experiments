package types

// Balance represents an L1 token balance
type Balance struct {
	Network      [4]byte `ssz-size:"4"`
	TokenAddress []byte  `ssz-max:"128"`
	// Amount in VGBit
	Amount uint64
	// Locked is the amount of tokens locked, immobile. Always locked <= amount
	Locked uint64
	// Delegated is the amount delegated to validator, always delegated <= locked
	Delegated uint64
}
