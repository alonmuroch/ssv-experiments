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

func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p := ssv2.NewPipeline(test.Pre)
	for _, msg := range test.Messages {
		err, _ := p.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: []spec_test.TestObject{test.Pre},
		Actual:         []spec_test.TestObject{test.Post},
	}
}
