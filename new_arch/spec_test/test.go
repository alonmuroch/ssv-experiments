package spec_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// SpecTest is a generic test struct runing and verifyin all tests with the goal of maximizing code and standard sharing between tests
type SpecTest[T TestImpl] struct {
	Test T
}

func NewSpecTest[T TestImpl](testImpl T) (*SpecTest[T], error) {
	ret := &SpecTest[T]{
		Test: testImpl,
	}
	return ret, nil
}

func (test *SpecTest[T]) Run(t *testing.T) {
	result := test.Test.Test(t)

	require.True(t, len(result.ExpectedResult) == len(result.Actual))
	for i := range result.ExpectedResult {
		expected := result.ExpectedResult[i]
		actual := result.Actual[i]

		expectedR, err := expected.HashTreeRoot()
		require.NoError(t, err)
		actualR, err := actual.HashTreeRoot()
		require.NoError(t, err)

		require.EqualValues(t, expectedR, actualR)
	}

	if result.BroadcastedBeaconObjects.NotEmpty() {
		// TODO test broadcasted
	}

	if result.BroadcastedMessages.NotEmpty() {
		// TODO test broadcasted
	}
}
