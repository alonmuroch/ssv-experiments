package runner

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/p2p"
	ssv2 "ssv-experiments/new_arch/pipeline/ssv"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
	"testing"
)

type SpecTest struct {
	Pre      *ssv.Runner
	Post     *ssv.Runner
	Share    *types.Share
	Duty     *types.Duty
	Messages []*p2p.Message `ssz-max:"256"`
}

func (test *SpecTest) Init(testSSZ []byte) (spec_test.TestObject, error) {
	if err := test.UnmarshalSSZ(testSSZ); err != nil {
		return nil, err
	}

	return test.Post, nil
}

func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p := ssv2.NewPipeline(test.Pre)
	for _, msg := range test.Messages {
		err, _ := p.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: test.Pre,
	}
}
