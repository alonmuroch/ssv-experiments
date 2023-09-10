package full_flow

import (
	"ssv-experiments/new_arch_2/tests/spec/asgard/qbft"
)

// FullFlow2
// @generate-test
func FullFlow2() *qbft.ProcessMessageTest {
	return nil
	//msgs := []*types.QBFTSignedMessage{
	//	fixtures.QBFTSignedMessage(1, types.FirstRound, types.ProposalMessageType),
	//	fixtures.QBFTSignedMessage(1, types.FirstRound, types.PrepareMessageType),
	//	fixtures.QBFTSignedMessage(2, types.FirstRound, types.PrepareMessageType),
	//	fixtures.QBFTSignedMessage(3, types.FirstRound, types.PrepareMessageType),
	//	fixtures.QBFTSignedMessage(1, types.FirstRound, types.CommitMessageType),
	//	fixtures.QBFTSignedMessage(2, types.FirstRound, types.CommitMessageType),
	//	fixtures.QBFTSignedMessage(3, types.FirstRound, types.CommitMessageType),
	//}
	//
	//return &qbft.ProcessMessageTest{
	//	Pre: &types.QBFT{
	//		Round:    types.FirstRound,
	//		Height:   fixtures.Height,
	//		Messages: []*types.QBFTSignedMessage{},
	//
	//		StartValue: fixtures.AttesterConsensusData,
	//	},
	//	Post: &types.QBFT{
	//		Round:    types.FirstRound,
	//		Height:   fixtures.Height,
	//		Messages: msgs,
	//
	//		PreparedRound:                   types.FirstRound,
	//		ProposalAcceptedForCurrentRound: fixtures.QBFTSignedMessage(1, types.FirstRound, types.ProposalMessageType),
	//
	//		StartValue: fixtures.AttesterConsensusData,
	//	},
	//	Messages: msgs,
	//}
}
