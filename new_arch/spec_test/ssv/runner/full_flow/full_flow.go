package full_flow

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test/fixtures"
	"ssv-experiments/new_arch/spec_test/ssv/runner"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

func FullFlow() *runner.SpecTest {
	msg := []*p2p.Message{
		fixtures.P2PQBFTSignedMessage(1, qbft.FirstRound, qbft.ProposalMessageType),
		fixtures.P2PQBFTSignedMessage(1, qbft.FirstRound, qbft.PrepareMessageType),
		fixtures.P2PQBFTSignedMessage(2, qbft.FirstRound, qbft.PrepareMessageType),
		fixtures.P2PQBFTSignedMessage(3, qbft.FirstRound, qbft.PrepareMessageType),
		fixtures.P2PQBFTSignedMessage(1, qbft.FirstRound, qbft.CommitMessageType),
		fixtures.P2PQBFTSignedMessage(2, qbft.FirstRound, qbft.CommitMessageType),
		fixtures.P2PQBFTSignedMessage(3, qbft.FirstRound, qbft.CommitMessageType),

		fixtures.P2PPartialSignatureMessage(1, fixtures.Slot, types.PostConsensusPartialSig),
		fixtures.P2PPartialSignatureMessage(2, fixtures.Slot, types.PostConsensusPartialSig),
		fixtures.P2PPartialSignatureMessage(3, fixtures.Slot, types.PostConsensusPartialSig),
	}

	cdByts, err := fixtures.AttesterConsensusData.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}

	return &runner.SpecTest{
		Role: types.BeaconRoleAttester,
		Pre: &ssv.Runner{
			State: &ssv.State{
				PartialSignatures: ssv.Container{},
				StartingDuty:      fixtures.AttesterDuty,
			},
			Share:      fixtures.Share,
			Identifier: fixtures.Identifier,
		},
		Post: &ssv.Runner{
			State: &ssv.State{
				PartialSignatures: ssv.Container{
					fixtures.PartialSignatureMessage(1, fixtures.Slot, types.PostConsensusPartialSig),
					fixtures.PartialSignatureMessage(2, fixtures.Slot, types.PostConsensusPartialSig),
					fixtures.PartialSignatureMessage(3, fixtures.Slot, types.PostConsensusPartialSig),
				},
				DecidedValue: cdByts,
				StartingDuty: fixtures.AttesterDuty,
			},
			Share:      fixtures.Share,
			Identifier: fixtures.Identifier,
		},
		Messages: msg,
	}
}
