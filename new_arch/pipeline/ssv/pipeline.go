package ssv

import (
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
)

func NewPipeline(runner *ssv.Runner) *pipeline.Pipeline {
	ret := pipeline.NewPipeline()
	ret.Runner = runner
	return ret
}
