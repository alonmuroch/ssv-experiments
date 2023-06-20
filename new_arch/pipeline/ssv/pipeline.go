package ssv

import (
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/ssv"
)

func NewPipeline(runner *ssv.Runner) *pipeline.Pipeline {
	return &pipeline.Pipeline{
		Runner: runner,
		Items:  []pipeline.PipelineF{},
		Phase:  map[string]int{},
	}
}
