package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

const FirstRound = 1

func Start(
	state *types.QBFT,
	share *types.Share,
	startValue *types.ConsensusData,
	height uint64,
) (*types.QBFTMessage, error) {
	state.StartValue = startValue
	state.Round = FirstRound
	state.Height = height

	if proposerForRound(FirstRound) == share.OperatorID {
		return CreateProposalMessage(state)
	}
	return nil, nil
}
