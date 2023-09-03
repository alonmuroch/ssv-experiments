package qbft

import (
	"github.com/pkg/errors"
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

const FirstRound = 1

type Instance struct {
	State      *types.QBFT
	Share      *types.Share
	StartValue *types.ConsensusData
}

func NewInstance(data *types.ConsensusData, share *types.Share, height, role uint64) *Instance {
	return &Instance{
		State: &types.QBFT{
			Round:  FirstRound,
			Height: height,
		},
		Share:      share,
		StartValue: data,
	}
}

func ProcessMessage(state *types.QBFT, message *types.QBFTSignedMessage) error {
	panic("implement")
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
func (i *Instance) ProcessMessage(msg *types.QBFTSignedMessage) (*types.QBFTSignedMessage, error) {
	// TODO process
	return nil, nil
}

// Decided returns true if decided.
func (i *Instance) Decided() bool {
	found, _, _ := i.DecidedRoot()
	return found
}

// DecidedRoot returns the root and messages that decided current round
func (i *Instance) DecidedRoot() (bool, []*types.QBFTSignedMessage, [32]byte) {
	byRoot := make(map[[32]byte][]*types.QBFTSignedMessage)

	// batch messages by root. If exists a decided message return immediately
	for _, m := range RoundAndType(i.State, i.State.Round, types.CommitMessageType) {
		// decided message return true
		if len(m.Signers) >= int(i.Share.Quorum) {
			return true, []*types.QBFTSignedMessage{m}, m.Message.Root
		}

		if byRoot[m.Message.Root] == nil {
			byRoot[m.Message.Root] = []*types.QBFTSignedMessage{}
		}
		byRoot[m.Message.Root] = append(byRoot[m.Message.Root], m)
	}

	// find if decided
	for _, msgs := range byRoot {
		if len(msgs) >= int(i.Share.Quorum) {
			return true, msgs, msgs[0].Message.Root
		}
	}
	return false, nil, [32]byte{}
}

// DecidedValue returns decided value for current round
func (i *Instance) DecidedValue() ([]byte, error) {
	decided, msgs, _ := i.DecidedRoot()
	if !decided {
		return nil, errors.New("not decided")
	}

	// single decided message
	if len(msgs) == 1 {
		return msgs[0].FullData, nil
	}

	// regular commit quorum
	proposalMsgs := RoundAndType(i.State, i.State.Round, types.ProposalMessageType)
	if len(proposalMsgs) != 1 {
		return nil, errors.New("no valid proposal for round")
	}
	return proposalMsgs[0].FullData, nil
}
