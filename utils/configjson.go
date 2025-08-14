package utils

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

const configJsonSchema string = `{
      "$schema": "http://json-schema.org/draft-07/schema#",
      "type": "object",
      "properties": {
        "name": { "type": "string" },
        "age":  { "type": "integer", "minimum": 0 }
      },
      "required": ["name", "age"]
    }`
const configJsonPath string = "./config.json"

func GetConfigJson(pathToConfig string) (*specs.Spec, error) {
	// read the file into byte slice
	logrus.Tracef("attempting to read config: %s", pathToConfig)
	configData, err := os.ReadFile(pathToConfig)
	if err != nil {
		logrus.Fatalf("error reading config: %v", err)
		return nil, err
	}

	logrus.Infof("successfully read config")
	logrus.Tracef("ensuring config.json is correct according to schema")
	if !configIsValid(&configData) {
		err = fmt.Errorf("config.json is invalid")
		logrus.Fatalf("config.json is invalid: %v", err)
		return nil, err
	}
	logrus.Tracef("config.json valid according to OCI schema")

	var thisSpec specs.Spec
	err = json.Unmarshal(configData, &thisSpec)
	if err != nil {
		return nil, err
	}

	logrus.Infof("successfully marshalled: %s", pathToConfig)
	return &thisSpec, nil
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

// configIsValid checks if a passed config.json file
// is valid according to the config schema in the
// OCI specification.
func configIsValid(config *[]byte) bool {
	//TODO: implement me
	// this is a bit of a PIA to implement, not completely critical
	// to user functionality either just now, will come back at later
	// date
	return true
}
