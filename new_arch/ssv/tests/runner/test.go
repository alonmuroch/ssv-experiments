package runner

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
	"testing"
)

type SpecTest struct {
	Pre      *ssv.State
	Post     *ssv.State
	Share    *types.Share
	Duty     *types.Duty
	Messages []*p2p.Message `ssz-max:"256"`

	runner *ssv.Runner
}

func (test *SpecTest) Init(testSSZ []byte) (spec_test.TestObject, error) {
	if err := test.UnmarshalSSZ(testSSZ); err != nil {
		return nil, err
	}

	test.runner = ssv.NewRunner(ssv.Config{
		Share: test.Share,
	}, test.Duty)

	return test.Post, nil
}

func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	for _, msg := range test.Messages {
		require.NoError(t, test.runner.ProcessMessage(msg))
	}

	return &spec_test.TestResult{
		ExpectedResult: &test.runner.State,
	}
}
