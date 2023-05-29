package types

type Duty struct {
	Role        uint64
	ValidatorPK [48]byte `ssz-size:"48"`
	Slot        uint64
}
