package src

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

type ParseConfigLink struct {
	next ChainLink
}

func (pc *ParseConfigLink) Execute(spec *specs.Spec) {
	panic("implement me")
}

func (pc *ParseConfigLink) SetNext(next ChainLink) {
	pc.next = next
}

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

func GetSpec() (*specs.Spec, error) {
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
