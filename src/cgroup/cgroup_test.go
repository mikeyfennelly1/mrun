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

	mjr := int64(12)
	mnr := int64(10)
	r.Devices = []specs.LinuxDeviceCgroup{
		{
			Allow:  false,
			Type:   "block",
			Major:  &mjr,
			Minor:  &mnr,
			Access: "rwm",
		},
	}

	err = InitCgroup("test-cgroup.slice", r)
	require.NoError(t, err)
}
