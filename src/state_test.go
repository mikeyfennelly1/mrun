package src

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"testing"
)

const containerID = "containerID"

func TestContainerManager_CreateAndInitStateFile(t *testing.T) {
	m := GetContainerManager(containerID)
	err := m.CreateAndInitStateFile()
	require.NoError(t, err)
}

func TestContainerManager_UpdateContainerStateFile(t *testing.T) {
	testingContainerStateObj := specs.State{
		Version:     ociVersion,
		ID:          containerID,
		Status:      "running",
		Pid:         2000,
		Bundle:      "/some/path",
		Annotations: map[string]string{},
	}
	m := GetContainerManager(containerID)

	err := m.UpdateContainerStateFile(testingContainerStateObj)
	require.NoError(t, err)
}

func TestContainerManager_DeleteStateFile(t *testing.T) {
	m := GetContainerManager(containerID)
	err := m.DeleteStateFile()
	require.NoError(t, err)
}
