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

// Init from an encoded ssz test, returns the expected post test object to be verified after test is run
func (test *SpecTest) Init(testSSZ []byte) (spec_test.TestObject, error) {
	if err := test.UnmarshalSSZ(testSSZ); err != nil {
		return nil, err
	}
	return test.Post, nil
}

// Test will run the test, fail if errors during test and will return a post run test object to be compared with
func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p := qbft2.NewPipeline(test.Pre)
	for _, msg := range test.Messages {
		err, _ := p.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: test.Pre,
	}
}
