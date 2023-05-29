package ssv

import "ssv-experiments/new_arch/types"

type PartialSignatureType = uint64

const (
	PreConsensus PartialSignatureType = iota
	PostConsensus
)

// PartialSignatureContainer holds partial BLS signature messages for various types
type PartialSignatureContainer struct {
	Types      []PartialSignatureType                  `ssz-max:"24"`
	Containers []*types.SignedPartialSignatureMessages `ssz-max:"24"`
}

func (c *PartialSignatureContainer) Add(t PartialSignatureType, msg *types.SignedPartialSignatureMessages) {

}
