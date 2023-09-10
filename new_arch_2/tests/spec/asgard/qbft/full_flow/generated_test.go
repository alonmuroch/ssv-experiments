package full_flow

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/qbft"
	"testing"
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
