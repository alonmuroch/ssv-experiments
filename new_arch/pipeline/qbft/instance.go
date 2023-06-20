package qbft

import (
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewAttesterPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	return NewPipeline(instance)

}
