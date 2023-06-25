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
	StartValue *types.ConsensusData
}

func NewInstance(data *types.ConsensusData, share *types.Share, height, role uint64) *Instance {
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

// Decided returns true if decided.
func (i *Instance) Decided() bool {
	found, _ := i.DecidedRoot()
	return found
}

func (i *Instance) DecidedRoot() (bool, [32]byte) {
	byRoot := make(map[[32]byte][]*SignedMessage)

	// batch messages by root. If exists a decided message return immediately
	for _, m := range i.State.Messages {
		if m.Message.MsgType != CommitMessageType {
			continue
		}

		// if decided message return true
		if len(m.Signers) >= int(i.Share.Quorum) {
			return true, m.Message.Root
		}

		if byRoot[m.Message.Root] == nil {
			byRoot[m.Message.Root] = []*SignedMessage{}
		}
		byRoot[m.Message.Root] = append(byRoot[m.Message.Root], m)
	}

	// find if decided
	for _, msgs := range byRoot {
		if len(msgs) >= int(i.Share.Quorum) {
			return true, msgs[0].Message.Root
		}
	}
	return false, [32]byte{}
}
