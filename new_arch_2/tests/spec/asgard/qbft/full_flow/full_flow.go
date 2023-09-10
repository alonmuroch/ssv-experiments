package full_flow

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/qbft"
)

func FullFlow() *qbft.ProcessMessageTest {
	msgs := []*types.QBFTSignedMessage{
		fixtures.QBFTSignedMessage(1, types.FirstRound, types.ProposalMessageType),
		fixtures.QBFTSignedMessage(1, types.FirstRound, types.PrepareMessageType),
		fixtures.QBFTSignedMessage(2, types.FirstRound, types.PrepareMessageType),
		fixtures.QBFTSignedMessage(3, types.FirstRound, types.PrepareMessageType),
		fixtures.QBFTSignedMessage(1, types.FirstRound, types.CommitMessageType),
		fixtures.QBFTSignedMessage(2, types.FirstRound, types.CommitMessageType),
		fixtures.QBFTSignedMessage(3, types.FirstRound, types.CommitMessageType),
	}

	return &qbft.ProcessMessageTest{
		Pre:      &types.QBFT{},
		Post:     &types.QBFT{},
		Messages: msgs,
	}
}
