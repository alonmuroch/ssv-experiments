package module

// Fork is a unique identifier for a module fork version, 2 identical pieces of data signed with different domains will result in different sigs
type Fork [4]byte

// Share holds cryptographic information about a cluster
type Share struct {
	OperatorID uint64
	Domain     Fork `ssz-size:"4"`
	Key        Key  `ssz-size:"4"`

	Cluster Cluster `ssz-max:"13"`

	Quorum        uint64
	PartialQuorum uint64

	// ExtraData holds share specific extra data needed for the module (e.g ETH staking module requires validator PK, fee recipient and graffiti)
	ExtraData []byte `ssz-max:"2048"`
}
