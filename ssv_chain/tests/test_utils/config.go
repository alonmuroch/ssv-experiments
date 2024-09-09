package test_utils

import "ssv-experiments/ssv_chain/types"

func (backend *TestingBackend) Config() *types.Configure {
	return &types.Configure{
		SupportedNetworks:        backend.SupportedNetworks,
		SystemTxSigner:           TestingSystemSigner,
		SSVTokenAddressByNetwork: backend.SSVTokenAddresses,
	}
}
