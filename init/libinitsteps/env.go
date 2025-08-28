package libinitsteps

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"strings"
)

type setEnvVarsStep struct{}

func (sev *setEnvVarsStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	panic("implement me")
	return nil
}

func setEnvVars(envVars []string) {
	for _, envVar := range envVars {
		setEnvVar(envVar)
	}
}

func setEnvVar(envVar string) {
	path := "PATH="
	if strings.HasPrefix(envVar, path) {
		pathVars := strings.Split(strings.TrimPrefix(envVar, path), ":")
		for _, e := range pathVars {

			logrus.Infof("path variable: %s\n", e)
		}
	}
}
