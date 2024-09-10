package test_utils

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/stretchr/testify/require"
	"testing"
)

func RequireEqualRoots(t *testing.T, a, b ssz.HashRoot) {
	aRoot, err := a.HashTreeRoot()
	require.NoError(t, err)
	bRoot, err := b.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, aRoot[:], bRoot[:])
}
