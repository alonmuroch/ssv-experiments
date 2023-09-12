package main

import (
	"flag"
)

const (
	GenerateTestsMode    = "generate-tests"
	GenerateAllTestsMode = "generate-all-tests"
	GenerateSSZMode      = "generate-ssz"
)

func main() {
	modePtr := flag.String("mode", GenerateTestsMode, "mode of operation")
	flag.Parse()

	switch *modePtr {
	case GenerateTestsMode:
		GenerateTests()
	case GenerateAllTestsMode:
		GenerateAllTestsFile()
	case GenerateSSZMode:
		GenerateSSZ()
	default:
		panic("mode not supported")
	}
}
