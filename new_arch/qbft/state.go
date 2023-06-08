package qbft

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

type State struct {
	Share      *types.Share
	Identifier p2p.Identifier `ssz-size:"56"` // instance Identifier this msg belongs to
	Height     uint64

	Messages Container
}
