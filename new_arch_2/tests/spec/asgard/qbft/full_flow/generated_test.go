package full_flow

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
	"testing"
	"ssv-experiments/new_arch_2/spec/asgard/types"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/qbft"
)

func TestFullFlow(t *testing.T) {
	tst, err := tests.NewSpecTest[*qbft.ProcessMessageTest](FullFlow())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}
func TestFullFlow2(t *testing.T) {
	tst, err := tests.NewSpecTest[*qbft.ProcessMessageTest](FullFlow2())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}

var AllTests = []tests.TestObject{FullFlow(), FullFlow2()}
