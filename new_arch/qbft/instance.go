package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

type Instance struct {
	State        State
	Share        *types.Share
	Identifier   p2p.Identifier `ssz-size:"56"` // instance Identifier this msg belongs to
	StartValue   InputData
	DecidedValue *InputData
}

func NewInstance(data *InputData, share *types.Share, height, role uint64) *Instance {
	return &Instance{
		State: State{
			Height: height,
		},
		Share:      share,
		Identifier: p2p.NewIdentifier(height, share.ValidatorPubKey, role),
		StartValue: *data,
	}
}

func (i *Instance[T]) StartForSlot(slot uint64) {

}

func (i *Instance[T]) ProcessMessage(msg *SignedMessage) error {
	if !bytes.Equal(msg.Message.Identifier[:], i.Identifier[:]) {
		return errors.New("invalid identifier")
	}
	// TODO process
	return nil
}

func (i *Instance[T]) Decided() bool {
	panic("implement")
}
