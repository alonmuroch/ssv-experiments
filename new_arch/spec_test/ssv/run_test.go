package ssv

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/spec_test/ssv/runner"
	full_flow2 "ssv-experiments/new_arch/spec_test/ssv/runner/full_flow"
	"testing"
)

func TestSpecTest(t *testing.T) {
	test, err := spec_test.NewSpecTest[*runner.SpecTest](full_flow2.FullFlow())
	require.NoError(t, err)

	test.Run(t)
}
