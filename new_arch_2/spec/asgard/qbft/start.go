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
	state.Round = types.FirstRound
	state.Height = height

	if IsProposer(state, share) {
		return CreateProposalMessage(state, startValue)
	}
	return nil, nil
}
