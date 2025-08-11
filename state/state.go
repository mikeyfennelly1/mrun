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
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	varRunMrun = "/var/run/mrun/"
	ociVersion = "0.2.0"
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type StateSubsytem interface {
	InitState() error
	GetState() *specs.State
	UpdateState() *specs.State
	DeleteState() *specs.State
}

var stateSingleton *specs.State = nil

func GetStateForContainer(containerID string) (*specs.State, error) {
	if stateSingleton == nil {
		state, err := initializeState(containerID)
		if err != nil {
			return nil, err
		}
		stateSingleton = state
	}
	return stateSingleton, nil
}

// initializeState creates a directory in state
// directory /var/run/mrun/<container-id> and a file
// state.json in that directory. i.e
//
// /var/run/mrun/<container-id>/state.json
func initializeState(containerID string) (*specs.State, error) {
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
		Status:      "running",
		Pid:         os.Getpid(),
		Bundle:      pwd,
		Annotations: nil,
	}
	err = m.CreateAndInitStateFile(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
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

type containerState struct {
	Name           string
	ID             string
	Command        string
	Status         string
	BundleLocation string
}

func getSubdirectories(root string) ([]string, error) {
	var subdirs []string

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, filepath.Join(root, entry.Name()))
		}
	}

	return subdirs, nil
}

func NewContainerID() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
