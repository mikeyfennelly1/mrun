package cgroup

import (
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestConfigureCgroups(t *testing.T) {
	readJSON, err := os.ReadFile("/home/mfennelly/config.json")
	require.NoError(t, err)

	var r specs.LinuxResources
	err = json.Unmarshal(readJSON, &r)
	require.NoError(t, err)

	err = ConfigureCgroups(r)
	require.NoError(t, err)
}
