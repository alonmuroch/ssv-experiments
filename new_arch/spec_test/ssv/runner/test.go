package runner

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/pipeline"
	ssvPipeline "ssv-experiments/new_arch/pipeline/ssv"
	"ssv-experiments/new_arch/spec_test"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
	"testing"
)

type SpecTest struct {
	Pre      *ssv.Runner
	Post     *ssv.Runner
	Role     uint64
	Messages []*p2p.Message `ssz-max:"256"`
}

func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p, err := test.getPipeline(t)
	require.NoError(t, err)

	for _, msg := range test.Messages {
		err, _ := p.ProcessMessage(msg)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		Actual:         []spec_test.TestObject{test.Pre},
		ExpectedResult: []spec_test.TestObject{test.Post},
	}
}
func (test *SpecTest) getPipeline(t *testing.T) (*pipeline.Pipeline, error) {
	switch test.Role {
	case types.BeaconRoleAttester:
		return ssvPipeline.NewAttesterPipeline(test.Pre)
	case types.BeaconRoleProposer:
		return ssvPipeline.NewProposerRunnerForDuty(test.Pre)
	default:
		return nil, errors.New("unsupported test runner type")
	}
}
