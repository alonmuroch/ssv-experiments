package process

import (
	"github.com/stretchr/testify/require"
	qbft2 "ssv-experiments/new_arch/pipeline/qbft"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test"
	"testing"
)

type SpecTest struct {
	Pre      *qbft.Instance
	Post     *qbft.Instance
	Messages []*qbft.SignedMessage `ssz-max:"256"`
}

// Test will run the test, fail if errors during test and will return a post run test object to be compared with
func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p, err := qbft2.NewQBFTPipelineFromInstance(test.Pre)
	require.NoError(t, err)
	for _, msg := range test.Messages {
		err, _ = p.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: []spec_test.TestObject{test.Post},
		Actual:         []spec_test.TestObject{test.Pre},
	}
}
