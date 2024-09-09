package add_operator

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/tests/test_utils"
	"ssv-experiments/ssv_chain/types"
)

func Generate(backend *test_utils.TestingBackend) []ssz.Marshaler {
	tx := &types.Transaction{
		Address: backend.ETHAccountsAddresses[0],
		Nonce:   0,
		Operations: []*types.Operation{
			{
				Type:          [4]byte{types.OP_User, types.OP_Operator, types.OP_Add, types.OP_V0},
				OperationData: backend.AddOperatorOperation(),
			},
		},
	}

	return []ssz.Marshaler{
		backend.BaseState(),
		tx,
	}
}
