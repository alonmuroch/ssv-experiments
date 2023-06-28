package main

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"ssv-experiments/msg_validation_benchmarking/strategies/naive"
)

func main() {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	benchmark := NewBenchmarker(naive.NewNaiveStrategy())
	benchmark.Run()
}
