package ssv

import (
	"ssv-experiments/new_arch_2/spec/asgard/runner"
	"ssv-experiments/new_arch_2/spec/asgard/types"
	"ssv-experiments/new_arch_2/tests"
)

type ProcessMessageTest struct {
	Pre      *types.State
	Post     *types.State
	Messages []*types.SignedPartialSignatureMessages `ssz-max:"256"`
}

// Run will run the test, fail if errors during test and will return a post run test object to be compared with
func (test *ProcessMessageTest) Run(share *types.Share) *tests.TestResult {
	var lastErr error
	for _, msg := range test.Messages {
		if err := runner.ProcessMessage(test.Pre, share, msg); err != nil {
			lastErr = err
		}
	}

	return &tests.TestResult{
		Actual:         tests.TestObjects{test.Pre},
		ExpectedResult: tests.TestObjects{test.Post},
		Error:          lastErr,
	}
}
