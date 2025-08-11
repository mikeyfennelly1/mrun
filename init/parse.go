package init

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/init/libinit"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

type parseConfigLink struct {
	next libinit.ExecutableInitStep
}

func (pc *parseConfigLink) Execute(spec *specs.Spec) error {
	panic("implement me")
	return nil
}

func (pc *parseConfigLink) SetNext(next libinit.ExecutableInitStep) {
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
