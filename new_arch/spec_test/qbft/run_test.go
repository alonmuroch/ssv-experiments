package qbft

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/spec_test/qbft/process"
	"ssv-experiments/new_arch/spec_test/qbft/process/full_flow"
	"testing"
)

func TestSpecTest(t *testing.T) {
	test, err := spec_test.NewSpecTest[*process.SpecTest](full_flow.FullFlow())
	require.NoError(t, err)

	test.Run(t)
}
