package full_flow

import (
	"ssv-experiments/new_arch/pipeline/ssv/tests/runner"
	"ssv-experiments/new_arch/spec_test/fixtures"
	"ssv-experiments/new_arch/ssv"
)

func FullFlow() *runner.SpecTest {
	return &runner.SpecTest{
		Pre: &ssv.Runner{
			State: &ssv.State{
				PartialSignatures: ssv.Container{},
			},
			Share:      fixtures.Share,
			Identifier: fixtures.Identifier,
		},
	}
}
