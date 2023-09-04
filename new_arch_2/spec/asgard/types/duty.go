package types

type Duty struct {
	Role        uint64
	ValidatorPK [48]byte `ssz-size:"48"`
	Slot        uint64
	DomainData  [32]byte `ssz-size:"32"`
}
