package types

type Share struct {
	OperatorID      uint64
	ValidatorPubKey [48]byte `ssz-size:"48"`
	Domain          [4]byte  `ssz-size:"4"`

	Quorum        uint64
	PartialQuorum uint64
}
