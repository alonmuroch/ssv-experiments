package types

type Share struct {
	OperatorID      uint64
	ValidatorPubKey [48]byte `ssz-size:"48"`
	Domain          Domain   `ssz-size:"4"`

	Cluster []*Signer `ssz-max:"13"`

	Quorum        uint64
	PartialQuorum uint64

	FeeRecipientAddress [20]byte `ssz-size:"20"`
	Graffiti            []byte   `ssz-size:"32"`
}
