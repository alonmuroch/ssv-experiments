package main

import (
	"fmt"
	"ssv-experiments/msg_validation_benchmarking/strategies"
	"time"
)

type IBenchmark interface {
	Init(
		network strategies.INetwork,
		msgValidation func([]byte),
	)
	Start()
	GetQueueStats() (pending, popped int, avgProcessTimeSec float64)
}

type Benchmarker[T IBenchmark] struct {
	Benchmark T
}

func NewBenchmarker[T IBenchmark](benchmark T) *Benchmarker[T] {
	network := NewNetwork(
		NewFuzzer(time.Millisecond*50, time.Millisecond*300),
		time.Second*3,
		1000,
	)
	benchmark.Init(network, MessageValidation(pk))

	return &Benchmarker[T]{
		Benchmark: benchmark,
	}
}

func (b *Benchmarker[T]) Run() {
	b.Benchmark.Start()
	b.printPost()
}

func (b *Benchmarker[T]) printPost() {
	pending, popped, avgProcessTimeSec := b.Benchmark.GetQueueStats()
	fmt.Printf("Post stats:\n"+
		"	pending size: %d\n"+
		"	popped cnt: %d\n"+
		"	Avg process time: %f\n", pending, popped, avgProcessTimeSec)
}
