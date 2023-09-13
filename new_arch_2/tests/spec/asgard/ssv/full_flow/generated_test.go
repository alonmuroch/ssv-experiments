package full_flow

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
	"testing"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/ssv"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func TestAttester(t *testing.T) {
	tst, err := tests.NewSpecTest[*ssv.ProcessMessageTest](Attester())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}
func TestProposer(t *testing.T) {
	tst, err := tests.NewSpecTest[*ssv.ProcessMessageTest](Proposer())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}
