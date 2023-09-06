package qbft

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func Start(
	state *types.QBFT,
	share *types.Share,
	startValue *types.ConsensusData,
	height uint64,
) (*types.QBFTMessage, error) {
	state.StartValue = startValue
	state.Round = types.FirstRound
	state.Height = height

	if proposerForRound(types.FirstRound) == share.OperatorID {
		return CreateProposalMessage(state)
	}
	return nil, nil
}
