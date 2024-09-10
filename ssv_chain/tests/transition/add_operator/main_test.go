package add_operator

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/stretchr/testify/require"
	"ssv-experiments/ssv_chain"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/tests"
	"ssv-experiments/ssv_chain/tests/test_utils"
	"ssv-experiments/ssv_chain/types"
	"testing"
)

func TestSpec(t *testing.T) {
	for _, be := range tests.Backends {
		require.NoError(t, tests.WriteFixture(be.Name, Generate(be)))

		preState := &types.State{}
		postState := &types.State{}
		tx := &types.Transaction{}
		require.NoError(t, tests.LoadFixture(be.Name, []ssz.Unmarshaler{preState, tx, postState}))
		require.EqualValues(t, [4]byte{0x0, 0x0, 0x0, 0x1}, preState.Domain)

		require.NoError(t, ssv_chain.ProcessTransaction(
			&operations.Context{
				Config:   be.Config(),
				State:    preState,
				Account:  preState.Accounts[0],
				GasPrice: tx.GasPrice,
			},
			tx,
		))

		test_utils.RequireEqualRoots(t, preState, postState)
	}

}
