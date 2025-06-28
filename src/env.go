package src

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"strings"
)

type SetEnvVarsLink struct {
	next ChainLink
}

func (sev SetEnvVarsLink) Execute(spec *specs.Spec) {
	panic("implement me")
}

func (sev SetEnvVarsLink) SetNext(next ChainLink) {
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
