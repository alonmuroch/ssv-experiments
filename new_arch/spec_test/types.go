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
	ExpectedResult           TestObject
	BroadcastedMessages      TestObjects
	BroadcastedBeaconObjects TestObjects
}

type TestImpl interface {
	// Init from an encoded ssz test, returns the expected post test object to be verified after test is run
	Init(testSSZ []byte) (TestObject, error)
	// Test will run the test, fail if errors during test and will return a post run test object to be compared with
	Test(t *testing.T) *TestResult
}

type Network interface {
}

type Storage interface {
}

type Beacon interface {
}
