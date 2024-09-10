package add_operator

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/tests/test_utils"
	"ssv-experiments/ssv_chain/types"
)

func Generate(backend *test_utils.TestingBackend) []ssz.Marshaler {
	preState := backend.BaseState()

	tx := &types.Transaction{
		Address:  backend.ETHAccountsAddresses[0],
		Nonce:    0,
		GasPrice: 1,
		Operations: []*types.Operation{
			{
				Type:          [4]byte{types.OP_User, types.OP_Operator, types.OP_Add, types.OP_V0},
				OperationData: backend.AddOperatorOperation(),
			},
		},
	}

	postState := backend.BaseState()
	postState.Operators = append(postState.Operators, &types.Operator{
		Address:   preState.Accounts[0].Address,
		ID:        1,
		PublicKey: backend.OperatorKey1(),
		Module:    0,
		Tiers:     []*types.PriceTier{backend.SSVTokenOperatorPriceTier()},
	})
	postState.Accounts[0].Balances[0].Amount = 9985 * common.VGBitTenthSSV

	return []ssz.Marshaler{
		preState,
		tx,
		postState,
	}
}
