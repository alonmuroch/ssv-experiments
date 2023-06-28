package spec_test

import (
	ssz "github.com/ferranbt/fastssz"
	"testing"
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
}

type TestImpl interface {
	ssz.Marshaler
	// Test will run the test, returns the expected and actual test results
	Test(t *testing.T) *TestResult
}

type Network interface {
}

type Storage interface {
}

type Beacon interface {
}
