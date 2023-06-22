package tests

import (
	"ssv-experiments/new_arch/qbft/tests/process/full_flow"
	"testing"
)

func TestSpecTest(t *testing.T) {
	test := full_flow.FullFlow()

	test.Test(t)
}