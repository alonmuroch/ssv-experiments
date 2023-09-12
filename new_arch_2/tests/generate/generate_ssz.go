package main

import (
	"os"
	"path/filepath"
	"reflect"
	"ssv-experiments/new_arch_2/tests"
	"strings"
)

func GenerateSSZ() {
	//for pkgName, tests := range AllTests {
	//	marshalTests(pkgName, tests)
	//}
}

func marshalTests(pkgName string, tests []tests.TestObject) {
	dirName, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	// create folders if missing
	folderPath := strings.Replace(pkgName, "_", "/", -1)
	path := filepath.Join(dirName, "new_arch_2", "tests", "generate", folderPath)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err.Error())
	}

	for _, test := range tests {
		name := reflect.TypeOf(test).Name()

		byts, err := test.MarshalSSZ()
		if err != nil {
			panic(err.Error())
		}

		// write
		// You can also write it to a file as a whole.
		err = os.WriteFile(name+".ssz", byts, 0644)
		if err != nil {
			panic(err.Error())
		}
	}
}
