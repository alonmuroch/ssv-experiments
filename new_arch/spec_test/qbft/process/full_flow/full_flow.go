package full_flow

import (
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test/fixtures"
	"ssv-experiments/new_arch/spec_test/qbft/process"
)

func FullFlow() *process.SpecTest {
	msg := []*qbft.SignedMessage{
		fixtures.QBFTSignedMessage(1, 1, qbft.ProposalMessageType),
		fixtures.QBFTSignedMessage(1, 1, qbft.PrepareMessageType),
		fixtures.QBFTSignedMessage(2, 1, qbft.PrepareMessageType),
		fixtures.QBFTSignedMessage(3, 1, qbft.PrepareMessageType),
		fixtures.QBFTSignedMessage(1, 1, qbft.CommitMessageType),
		fixtures.QBFTSignedMessage(2, 1, qbft.CommitMessageType),
		fixtures.QBFTSignedMessage(3, 1, qbft.CommitMessageType),
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
			StartValue: fixtures.AttesterConsensusData,
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
			StartValue: fixtures.AttesterConsensusData,
		},
		Messages: msg,
	}
}
