//go:generate mockgen -source=state.go -destination=../mocks/state.go -package=mocks

// global defines the global information about running containers on
// the system that mrun manages.
//
// This comes in the form of state.json files for each container.
// These state.json files are found at a location in the filesystem
// that is defined by convention.
//
// see docs about mrun standard conventions for more.

package state

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	varRunMrun = "/var/run/mrun/"
	ociVersion = "0.2.0"
)

// StateSubsytem is used as an interface to creating and
// managing a state for a containerID.
//
// It creates a state.json file, at the location /var/run/mrun/<containerId>/state.json
//
// For more information on state.json, see: https://github.com/opencontainers/runtime-spec/blob/main/schema/state-schema.json
type StateSubsytem interface {
	InitState(containerID string) error
	GetStateManager(containerID string) *StateManager
	UpdateState(containerID string, updatedState *specs.State) error
	DeleteState(containerID string) error
}

var stateSingleton *specs.State = nil

// StateManager allows an object to exist that exposes an interface
// that someone can use to update the object, and in turn
// the state for that corresponding container.
//
// StateManager can only ever provide the interface to 1 state file.
type StateManager struct {
	state *specs.State
}

// initializeState creates a directory in state
// directory /var/run/mrun/<container-id> and a file
// state.json in that directory. i.e
//
// /var/run/mrun/<container-id>/state.json
func initializeState(containerID string) (*StateManager, error) {
	if containerID == "" {
		return nil, fmt.Errorf("containerId is invalid")
	}
	containerDirname := fmt.Sprintf("%s%s", varRunMrun, containerID)

	err := os.MkdirAll(containerDirname, 0775)
	if err != nil {
		logrus.Errorf("error while creating directory for managing state of container: %v", err)
		return nil, err
	}

	m := getContainerManager(containerID)
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	state := specs.State{
		Version:     ociVersion,
		ID:          containerID,
		Status:      "creating",
		Pid:         os.Getpid(),
		Bundle:      pwd,
		Annotations: nil,
	}
	err = m.CreateAndInitStateFile(&state)
	if err != nil {
		return nil, err
	}

	return &StateManager{
		state: &state,
	}, nil
}

type containerManager struct {
	containerID string
}

func getContainerManager(containerID string) containerManager {
	return containerManager{
		containerID: containerID,
	}
}

func (c *containerManager) UpdateContainerStateFile(state specs.State) error {
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

func (c *containerManager) DeleteStateFile() error {
	return os.Remove(c.getContainerStateFileName())
}

func (c *containerManager) GetContainerState() (*specs.State, error) {
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

func (c *containerManager) getContainerStateFileName() string {
	return fmt.Sprintf("%s/state.json", c.getContainerDirectoryName())
}

func (c *containerManager) getContainerDirectoryName() string {
	return fmt.Sprintf("%s%s", varRunMrun, c.containerID)
}

func (c *containerManager) CreateAndInitStateFile(state *specs.State) error {
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

	stateJSONData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.getContainerStateFileName(), stateJSONData, 0666)
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
