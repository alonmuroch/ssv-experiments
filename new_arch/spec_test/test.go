package spec_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type SpecTest[T TestImpl] struct {
	Test           T
	ExpectedResult TestObject
}

func NewSpecTest[T TestImpl](testImpl T, testSSZ []byte) (*SpecTest[T], error) {
	ret := &SpecTest[T]{
		Test: testImpl,
	}

	expected, err := ret.Test.Init(testSSZ)
	if err != nil {
		return nil, err
	}
	ret.ExpectedResult = expected
	return ret, nil
}

func (test *SpecTest[T]) Run(t *testing.T) {
	result := test.Test.Test(t)

	expectedR, err := test.ExpectedResult.HashTreeRoot()
	require.NoError(t, err)
	resultR, err := result.ExpectedResult.HashTreeRoot()
	require.EqualValues(t, expectedR, resultR)

	if result.BroadcastedBeaconObjects.NotEmpty() {
		// TODO test broadcasted
	}

	if result.BroadcastedMessages.NotEmpty() {
		// TODO test broadcasted
	}
}
