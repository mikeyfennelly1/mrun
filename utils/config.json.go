package utils

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

const configJsonPath string = "./config.json"

func parseConfig(pathToConfig string) (*specs.Spec, error) {
	configData, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, err
	}

	var thisSpec specs.Spec
	err = json.Unmarshal(configData, &thisSpec)
	if err != nil {
		return nil, err
	}

	return &thisSpec, nil
}

func getSpec() (*specs.Spec, error) {
	var spec specs.Spec
	jsonContent, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}
	err = json.Unmarshal(jsonContent, &spec)
	if err != nil {
		return nil, fmt.Errorf("error creating unmarshalling JSON: %v", err)
	}
	return &spec, nil
}

// ConfigJsonExists uses Unix stat syscall to check if file
// metadata is retrievable
func ConfigJsonExists() bool {
	// fail if stat information is not found
	s, err := os.Stat(configJsonPath)
	// if there is an error, fail
	if err != nil {
		return false
	}

	if s != nil {
		return true
	}
	return false
}

// GetConfigJsonContents returns the bytearray contents
// of $(pwd)/config.json.
//
// This should only really be used after ConfigJsonExists
// It does not care if there are no contents in the file
func GetConfigJsonContents() (*[]byte, error) {
	fileContents, _ := os.ReadFile(configJsonPath)
	if fileContents != nil {
		return nil, fmt.Errorf("failed to read %s", configJsonPath)
	}

	return &fileContents, nil
}

// configIsValid checks if a passed config.json file
// is valid according to the config schema in the
// OCI specification.
func configIsValid(config *string) bool {
	//TODO: implement me
	panic("configIsValid: implement me")
	return false
}
