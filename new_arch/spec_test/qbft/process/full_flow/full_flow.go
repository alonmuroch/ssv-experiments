package full_flow

import (
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test/fixtures"
	"ssv-experiments/new_arch/spec_test/qbft/process"
)

func FullFlow() *process.SpecTest {
	msg := []*qbft.SignedMessage{
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.ProposalMessageType},
			Signers: []uint64{1},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.PrepareMessageType},
			Signers: []uint64{1},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.PrepareMessageType},
			Signers: []uint64{2},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.PrepareMessageType},
			Signers: []uint64{3},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.CommitMessageType},
			Signers: []uint64{1},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.CommitMessageType},
			Signers: []uint64{2},
		},
		{
			Message: qbft.Message{Round: 1, MsgType: qbft.CommitMessageType},
			Signers: []uint64{3},
		},
	}

	return &process.SpecTest{
		Pre: &qbft.Instance{
			State: &qbft.State{
				Round:    qbft.FirstRound,
				Height:   0,
				Messages: qbft.NewContainer(),
			},
			Share:      fixtures.Share,
			Identifier: fixtures.Identifier,
			StartValue: fixtures.InputData,
		},
		Post: &qbft.Instance{
			State: &qbft.State{
				Round:         qbft.FirstRound,
				Height:        0,
				Messages:      msg,
				PreparedRound: 1,
			},
			Share:      fixtures.Share,
			Identifier: fixtures.Identifier,
			StartValue: fixtures.InputData,
		},
		Messages: msg,
	}
}
