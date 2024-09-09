package test_utils

import (
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/types"
)

var (
	TestingSystemSigner = make([]byte, 128)
)

func (backend *TestingBackend) BaseState() *types.State {
	return &types.State{
		Domain:                backend.Domain,
		LatestBlockHeaderHash: make([]byte, 32),
		Modules: []*types.Module{
			{
				Network: _to4Bytes(backend.SupportedNetworks[0]),
				Address: backend.ETHAccountsAddresses[0],
				Name:    []byte("test_module"),
			},
		},
		Accounts: []*types.Account{
			{
				Network: _to4Bytes(backend.SupportedNetworks[0]),
				Address: backend.ETHAccountsAddresses[0],
				Balances: []*types.Balance{
					{
						Network:      _to4Bytes(backend.SupportedNetworks[0]),
						TokenAddress: backend.SSVTokenAddresses[0],
						Amount:       common.OneSSV,
					},
				},
			},
		},
		Operators: []*types.Operator{
			{
				Address:   backend.ETHAccountsAddresses[0],
				PublicKey: backend.OperatorKey1(),
				Module:    0,
			},
		},
	}
}
