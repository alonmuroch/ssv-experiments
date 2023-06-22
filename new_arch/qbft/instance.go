package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

const FirstRound = 1

type Instance struct {
	State      *State
	Share      *types.Share
	Identifier p2p.Identifier `ssz-size:"56"`
	StartValue *InputData
}

func NewInstance(data *InputData, share *types.Share, height, role uint64) *Instance {
	return &Instance{
		State: &State{
			Height: height,
		},
		Share:      share,
		Identifier: p2p.NewIdentifier(height, share.ValidatorPubKey, role),
		StartValue: data,
	}
}

func (i *Instance) IsFirstRound() bool {
	return i.State.Round == FirstRound
}

// IsProposer returns true if propsoer for current round
func (i *Instance) IsProposer() bool {
	return i.proposerForRound(i.State.Round) == i.Share.OperatorID
}

func (i *Instance) proposerForRound(round uint64) uint64 {
	// TODO round robin
	return 1
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
