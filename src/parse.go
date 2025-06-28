package src

import (
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

type ParseConfigLink struct {
	next ChainItem
}

func (pc *ParseConfigLink) Execute(spec *specs.Spec) {
	panic("implement me")
}

func (pc *ParseConfigLink) SetNext(next ChainItem) {
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
