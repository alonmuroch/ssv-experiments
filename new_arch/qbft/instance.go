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

// Start will start the instance and return the messages to broadcast
func (i *Instance) Start() (*SignedMessage, error) {
	if i.proposerForRound(i.State.Round) == i.Share.OperatorID {
		return i.createProposal()
	}
	return nil, nil
}

func (i *Instance) proposerForRound(round uint64) uint64 {
	panic("implement")
}

// ProcessMessage processes the incoming message and returns an optional message to be broadcasted. Or error
func (i *Instance) ProcessMessage(msg *SignedMessage) (*SignedMessage, error) {
	if !bytes.Equal(msg.Message.Identifier[:], i.Identifier[:]) {
		return nil, errors.New("invalid identifier")
	}
	// TODO process
	return nil, nil
}

func (i *Instance) Decided() bool {
	panic("implement")
}
