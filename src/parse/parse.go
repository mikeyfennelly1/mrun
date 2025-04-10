package parse

import (
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

func ParseConfig(pathToConfig string) (*specs.Spec, error) {
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
