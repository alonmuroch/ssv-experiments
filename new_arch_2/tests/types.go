package tests

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

type TestObject interface {
	ssz.Marshaler
	ssz.HashRoot
}

type TestObjects []TestObject

func (obj TestObjects) NotEmpty() bool {
	return len(obj) > 0
}

type TestResult struct {
	ExpectedResult           TestObjects
	Actual                   TestObjects
	BroadcastedMessages      TestObjects
	BroadcastedBeaconObjects TestObjects
	Error                    error
}

type TestImpl interface {
	ssz.Marshaler
	// Run will run the test, returns the expected and actual test results
	Run(share *types.Share) *TestResult
}
