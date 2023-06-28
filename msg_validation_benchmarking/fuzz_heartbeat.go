package main

import (
	"math/rand"
	"time"
)

type FuzzerHeartBeat struct {
	Upper, Lower time.Duration
}

func NewFuzzer(upper, lower time.Duration) *FuzzerHeartBeat {
	return &FuzzerHeartBeat{
		Upper: upper,
		Lower: lower,
	}
}

func (f *FuzzerHeartBeat) RunWithFuzzDelay(fnc func()) {
	go func(fnc func()) {
		time.Sleep(randTimeout(f.Lower, f.Upper))
		fnc()
	}(fnc)
}

func randTimeout(lower, upper time.Duration) time.Duration {
	rand.Seed(time.Now().UnixNano())
	n := int(lower) + rand.Intn(int(upper-lower+1))
	return time.Duration(n)
}
