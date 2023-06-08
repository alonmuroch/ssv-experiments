package ssv

import "ssv-experiments/new_arch/types"

type Container []*types.SignedPartialSignatureMessages

func (c Container) AllPreConsensus() []*types.SignedPartialSignatureMessages {
	panic("implement")
}

func (c Container) AllPostConsensus() []*types.SignedPartialSignatureMessages {
	panic("implement")
}
