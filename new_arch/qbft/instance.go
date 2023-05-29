package qbft

import (
	"bytes"
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

type ProposedValueCheckF[T IData] func(data T) error

type IData interface {
	ssz.Marshaler | ssz.HashRoot
}

type Instance[T IData] struct {
	State        State
	StartValue   T
	DecidedValue T

	valueCheckF ProposedValueCheckF[T]
}

func NewInstance[T IData](data T, valueCheckF ProposedValueCheckF[T], share *types.Share, slot, role uint64) *Instance[T] {
	return &Instance[T]{
		State: State{
			Share:      share,
			Identifier: p2p.NewIdentifier(slot, share.ValidatorPubKey, role),
			Height:     slot,
		},
		StartValue: data,

		valueCheckF: valueCheckF,
	}
}

func (i *Instance[T]) StartForSlot(slot uint64) {

}

func (i *Instance[T]) ProcessMessage(msg *SignedMessage) error {
	if !bytes.Equal(msg.Message.Identifier[:], i.State.Identifier[:]) {
		return errors.New("invalid identifier")
	}
	// TODO process
	return nil
}

func (i *Instance[T]) Decided() bool {
	panic("implement")
}
