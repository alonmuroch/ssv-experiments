package tests

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	ssz "github.com/ferranbt/fastssz"
	"os"
	"path/filepath"
)

const (
	FixtureFilename          = "fixture.ssz"
	FixtureGeneratorFilename = "test_gen.go"
)

func LoadFixture(prefix string, objs []ssz.Unmarshaler) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	fixturePath := filepath.Join(pwd, prefix+"_"+FixtureFilename)

	_, err = os.Stat(fixturePath)
	if os.IsNotExist(err) {
		return err
	}
	byts, err := os.ReadFile(fixturePath)
	if err != nil {
		return err
	}

	// unmarshal json
	encoded := map[int]string{}
	if err := json.Unmarshal(byts, &encoded); err != nil {
		return err
	}

	//decode ssz
	for i, obj := range encoded {
		byts, err := hex.DecodeString(obj)
		if err != nil {
			return err
		}
		if err := objs[i].UnmarshalSSZ(byts); err != nil {
			return err
		}
	}
	return nil
}

func WriteFixture(prefix string, objs []ssz.Marshaler) error {
	encoded := map[int]string{}
	for i, obj := range objs {
		byts, err := obj.MarshalSSZ()
		if err != nil {
			return err
		}
		encoded[i] = hex.EncodeToString(byts)
	}

	byts, err := json.Marshal(encoded)
	if err != nil {
		return err
	}
	return os.WriteFile(prefix+"_"+FixtureFilename, byts, 0644)
}
