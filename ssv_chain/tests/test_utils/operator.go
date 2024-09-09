package test_utils

import (
	"crypto/x509"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations/operator"
	"ssv-experiments/ssv_chain/types"
)

func (backend *TestingBackend) OperatorKey1() *common.CryptoKey {
	pk, _ := x509.MarshalPKIXPublicKey(&backend.RSASKs[0].PublicKey)

	return &common.CryptoKey{
		Type: [2]byte{common.PublicKey, common.RSA},
		Key:  pk,
	}
}

func (backend *TestingBackend) AddOperatorOperation() []byte {
	ret, _ := (&operator.AddOperatorV0{
		PublicKey: backend.OperatorKey1(),
		ModuleID:  0,
		Tiers: []*types.PriceTier{
			{
				Network:             _to4Bytes(backend.SupportedNetworks[0]),
				Capacity:            500,
				Price:               common.OneHundredthSSV,
				PayableTokenAddress: backend.SSVTokenAddresses[0],
			},
		},
	}).MarshalSSZ()

	return ret
}
