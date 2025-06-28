package src

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func SetEnvVars(envVars []string) {
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
