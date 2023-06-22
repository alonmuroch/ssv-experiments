package tests

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/pipeline/qbft/tests/process"
	"ssv-experiments/new_arch/pipeline/qbft/tests/process/full_flow"
	"ssv-experiments/new_arch/spec_test"
	"testing"
)

func TestSpecTest(t *testing.T) {
	test, err := spec_test.NewSpecTest[*process.SpecTest](full_flow.FullFlow())
	require.NoError(t, err)

	test.Run(t)
}
