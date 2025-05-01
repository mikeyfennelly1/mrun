package cgroup

import (
	"encoding/json"
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const testSliceName = "test-container.slice"

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

	m, err := InitCgroup(testSliceName, r)
	require.NoError(t, err)

	err = m.Update(cgroup2.ToResources(&r))
	require.NoError(t, err)

	require.NoError(t, err)
}

func TestUpdatePids(t *testing.T) {
	// Load the testing control group
	m, err := cgroup2.LoadSystemd("/", testSliceName)
	if err != nil {
		logrus.Errorf("Failed to create manager for cgroup slice: %s: %v\n", testSliceName, err)
		return
	}

	// read the test json config into a []byte
	readJSON, err := os.ReadFile("/home/mfennelly/config.json")
	require.NoError(t, err)

	// unmarshal test json to specs.LinuxResources
	var r specs.LinuxResources
	err = json.Unmarshal(readJSON, &r)
	require.NoError(t, err)

	// set the pids value in the LinuxResources obj
	r.Pids = &specs.LinuxPids{
		Limit: 20,
	}

	err = m.Update(cgroup2.ToResources(&r))
	require.NoError(t, err)
}

func TestKillSystemd(t *testing.T) {
	m, err := cgroup2.LoadSystemd("/", testSliceName)
	if err != nil {
		logrus.Errorf("Failed to create manager for cgroup slice: %s: %v\n", testSliceName, err)
		return
	}

	err = m.Kill()
	require.NoError(t, err)
}

func TestDeleteCgroup(t *testing.T) {
	m, err := cgroup2.LoadSystemd("/", testSliceName)
	if err != nil {
		logrus.Errorf("Failed to create manager for cgroup slice: %s: %v\n", testSliceName, err)
		return
	}

	err = m.DeleteSystemd()
	require.NoError(t, err)
}

func cleanUp() {
	m, err := cgroup2.LoadSystemd("/", testSliceName)
	if err != nil {
		logrus.Errorf("Failed to create manager for cgroup slice: %s: %v\n", testSliceName, err)
		return
	}

	err = m.Delete()
	if err != nil {
		logrus.Errorf("Failed to delete cgroup slice: %s\n", testSliceName)
		return
	}
}
