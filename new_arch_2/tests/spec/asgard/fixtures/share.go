package fixtures

import "ssv-experiments/new_arch_2/spec/asgard/types"

var Share = &types.Share{
	OperatorID:      1,
	ValidatorPubKey: [48]byte{},
	Domain:          [4]byte{},
	Quorum:          3,
	PartialQuorum:   2,
}
