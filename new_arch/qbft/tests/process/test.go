package process

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/types"
	"testing"
)

type SpecTest struct {
	Pre      *qbft.State
	Post     *qbft.State
	Share    *types.Share
	Messages []*qbft.SignedMessage `ssz-max:"256"`

	instance *qbft.Instance
}

// Init from an encoded ssz test, returns the expected post test object to be verified after test is run
func (test *SpecTest) Init(testSSZ []byte) (spec_test.TestObject, error) {
	if err := test.UnmarshalSSZ(testSSZ); err != nil {
		return nil, err
	}

	test.instance = &qbft.Instance{
		State: *test.Pre,
		Share: test.Share,
	}

	return test.Post, nil
}

// Test will run the test, fail if errors during test and will return a post run test object to be compared with
func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	for _, msg := range test.Messages {
		_, err := test.instance.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: &test.instance.State,
	}
}
