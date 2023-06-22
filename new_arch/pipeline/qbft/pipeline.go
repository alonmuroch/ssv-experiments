package qbft

import (
	"ssv-experiments/new_arch/pipeline"
	"ssv-experiments/new_arch/qbft"
)

func NewPipeline(instance *qbft.Instance) *pipeline.Pipeline {
	ret := pipeline.NewPipeline()
	ret.Instance = instance
	return ret
}
