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
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	varRunMrun   = "/var/run/mrun/"
	stateDotJSON = "state.json"
	ociVersion   = "0.2.0"
	letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type InitContainerStateLink struct {
	next ChainLink
}

func (ics *InitContainerStateLink) Execute(spec *specs.Spec) {
	panic("implement me")
}

func (ics *InitContainerStateLink) SetNext(next ChainLink) {
	ics.next = next
}

// InitContainerStateDirAndFile creates a directory in state
// directory /var/run/mrun/<container-id> and a file
// state.json in that directory. i.e
//
// /var/run/mrun/<container-id>/state.json
func InitContainerStateDirAndFile(containerID string, spec specs.Spec) error {
	containerDirname := fmt.Sprintf("%s%s", varRunMrun, containerID)

	err := os.MkdirAll(containerDirname, 0775)
	if err != nil {
		logrus.Errorf("error while creating directory for managing state of container: %v", err)
		return err
	}

	m := GetContainerManager(containerID)
	pwd, err := os.Getwd()
	if err != nil {
		return err
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
	return fmt.Sprintf("%s%s", varRunMrun, c.containerID)
}

func (c *ContainerManager) CreateAndInitStateFile(state *specs.State) error {
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

type ContainerState struct {
	Name           string
	ID             string
	Command        string
	Status         string
	BundleLocation string
}

func GetStateOfAllContainers() (*[]ContainerState, error) {
	subDirs, err := getSubdirectories(varRunMrun)
	if err != nil {
		return nil, err
	}

	var containerInfo []ContainerState
	containerInfo = []ContainerState{}
	for _, dir := range subDirs {
		thisContainerState, _ := GetContainerInfoByContainerID(dir)
		containerInfo = append(containerInfo, *thisContainerState)
	}

	return &containerInfo, nil
}

func GetContainerInfoByContainerID(containerID string) (*ContainerState, error) {

	return nil, nil
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
