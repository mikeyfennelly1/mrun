package init

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"strings"
)

type setEnvVarsLink struct {
	next ExecutableInitStep
}

func (sev setEnvVarsLink) Execute(spec *specs.Spec) error {
	panic("implement me")
	return nil
}

func (sev setEnvVarsLink) SetNext(next ExecutableInitStep) {
	sev.next = next
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
