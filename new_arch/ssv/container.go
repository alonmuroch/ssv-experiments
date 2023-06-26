package ssv

import "ssv-experiments/new_arch/types"

type Container []*types.SignedPartialSignatureMessages

func (c Container) AllPreConsensus() []*types.SignedPartialSignatureMessages {
	ret := make([]*types.SignedPartialSignatureMessages, 0)
	for _, m := range c {
		if m.Message.Type.IsPreConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}

func (c Container) AllPostConsensus() []*types.SignedPartialSignatureMessages {
	ret := make([]*types.SignedPartialSignatureMessages, 0)
	for _, m := range c {
		if m.Message.Type.IsPostConsensusType() {
			ret = append(ret, m)
		}
	}
	return ret
}
