// global defines the global information about running containers on
// the system that mrun manages.
//
// This comes in the form of state.json files for each container.
// These state.json files are found at a location in the filesystem
// that is defined by convention.
//
// see docs about mrun standard conventions for more.

package src

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	stateFilesLocation = "/var/run/mrun/"
	stateDotJSON       = "state.json"
	ociVersion         = "0.2.0"
)

// InitContainerStateDirAndFile creates a directory in state
// directory /var/run/mrun/<container-id> and a file
// state.json in that directory. i.e
//
// /var/run/mrun/<container-id>/state.json
func InitContainerStateDirAndFile(containerID string) error {
	containerDirname := fmt.Sprintf("%s%s", stateFilesLocation, containerID)
	containerStateFileAbsPath := fmt.Sprintf("%s/%s", containerDirname, stateDotJSON)

	err := os.MkdirAll(containerDirname, 0775)
	if err != nil {
		logrus.Errorf("error while creating directory for managing state of container: %v", err)
		return err
	}

	file, err := os.Create(containerStateFileAbsPath)
	if err != nil {
		return err
	}

	err = file.Chown(0, 0)
	if err != nil {
		logrus.Errorf("error changing ownership on state.json for container: %v", err)
		return err
	}

	return nil
}

type ContainerManager struct {
	containerID string
}

func GetContainerManager(containerID string) ContainerManager {
	return ContainerManager{
		containerID: containerID,
	}
}

func (c *ContainerManager) UpdateContainerStateFile(state specs.State) error {
	stateByteArray, err := json.Marshal(state)
	if err != nil {
		logrus.Errorf("could not update container state: %v", err)
		return err
	}

	err = os.WriteFile(c.getContainerStateFileName(), stateByteArray, 0775)
	if err != nil {
		logrus.Errorf("could not update container state: %v", err)
		return err
	}

	return nil
}

func (c *ContainerManager) DeleteStateFile() error {
	return os.Remove(c.getContainerStateFileName())
}

func (c *ContainerManager) GetContainerState() (*specs.State, error) {
	stateContents, err := os.ReadFile(c.getContainerStateFileName())
	if err != nil {
		return nil, err
	}

	var state specs.State
	err = json.Unmarshal(stateContents, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func (c *ContainerManager) getContainerStateFileName() string {
	return fmt.Sprintf("%s/state.json", c.getContainerDirectoryName())
}

func (c *ContainerManager) getContainerDirectoryName() string {
	return fmt.Sprintf("%s%s", stateFilesLocation, c.containerID)
}

func (c *ContainerManager) CreateAndInitStateFile() error {
	// create the stateFileLocation directory
	err := os.MkdirAll(c.getContainerDirectoryName(), 0775)
	if err != nil {
		logrus.Fatalf("could not create directory %s: %v", c.getContainerDirectoryName(), err)
		return err
	}

	_, err = os.Create(c.getContainerStateFileName())
	if err != nil {
		logrus.Fatalf("could not create file %s: %v", c.getContainerStateFileName(), err)
		return err
	}

	err = os.WriteFile(c.getContainerStateFileName(), []byte("{}"), 0666)
	if err != nil {
		logrus.Fatalf("could not initialize file %s: %v", c.getContainerStateFileName(), err)
		return err
	}

	err = os.Chown(c.getContainerStateFileName(), 0, 0)
	if err != nil {
		logrus.Errorf("could not transfer ownership of %s to user root: %v", c.getContainerStateFileName(), err)
		return err
	}

	logrus.Infof("succesfull initialization of container state file at %s", c.getContainerStateFileName())

	return nil
}
