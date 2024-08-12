package ssv_chain

import "ssv-experiments/ssv_chain/types"

type State struct {
	/*  Chain */
	// Domain represents: byte[0] empty, byte[1] empty, byte[2] network ID, byte[3] fork ID
	Domain                [4]byte
	BlockHeight           uint64
	LatestBlockHeaderHash [32]byte
	Validators            []*types.Validator
	/*  Registry */
	Accounts  []*types.Account
	Clusters  []*types.Cluster
	Operators []*types.Operator
	Modules   []*types.Module
}

// ProcessTransactions processes and applies transactions on state, returns nil if valid
func (s *State) ProcessTransactions(txs []*Transaction) error {
	panic("implement")
}
