package full_flow

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/ssv"
)

// Attester
// @generate-test
func Attester() *ssv.ProcessMessageTest {
	msg := []*types.SignedPartialSignatureMessages{
		fixtures.SignedPartialSignatureMessage(1, fixtures.Slot, types.PostConsensusPartialSig),
		fixtures.SignedPartialSignatureMessage(2, fixtures.Slot, types.PostConsensusPartialSig),
		fixtures.SignedPartialSignatureMessage(3, fixtures.Slot, types.PostConsensusPartialSig),
	}

	qbft := &types.QBFT{
		Round:  types.FirstRound,
		Height: fixtures.Height,
		Messages: []*types.QBFTSignedMessage{
			fixtures.QBFTSignedMessage(1, types.FirstRound, types.ProposalMessageType),
			fixtures.QBFTSignedMessage(1, types.FirstRound, types.PrepareMessageType),
			fixtures.QBFTSignedMessage(2, types.FirstRound, types.PrepareMessageType),
			fixtures.QBFTSignedMessage(3, types.FirstRound, types.PrepareMessageType),
			fixtures.QBFTSignedMessage(1, types.FirstRound, types.CommitMessageType),
			fixtures.QBFTSignedMessage(2, types.FirstRound, types.CommitMessageType),
			fixtures.QBFTSignedMessage(3, types.FirstRound, types.CommitMessageType),
		},

		PreparedRound:                   types.FirstRound,
		ProposalAcceptedForCurrentRound: fixtures.QBFTSignedMessage(1, types.FirstRound, types.ProposalMessageType),
	}

	return &ssv.ProcessMessageTest{
		Pre: &types.State{
			PartialSignatures: []*types.SignedPartialSignatureMessages{},
			QBFT:              qbft,
			StartingDuty:      fixtures.AttesterDuty,
		},
		Post: &types.State{
			PartialSignatures: msg,
			QBFT:              qbft,
			StartingDuty:      fixtures.AttesterDuty,
		},
		Messages: msg,
	}
}
