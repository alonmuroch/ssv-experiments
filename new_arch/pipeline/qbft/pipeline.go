package qbft

import (
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	return &pipeline.Pipeline{
		Instance: instance,
		Items:    []pipeline.PipelineF{},
		Phase:    map[string]int{},
	}
}
