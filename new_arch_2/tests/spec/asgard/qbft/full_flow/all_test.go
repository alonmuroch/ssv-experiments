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
