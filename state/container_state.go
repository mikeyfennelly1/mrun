package state

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

// StateManager allows an object to exist that exposes an interface
// that someone can use to update the object, and in turn
// the state for that corresponding container.
//
// StateManager can only ever provide the interface to 1 state file.
type StateManager struct {
	containerID string
}

func (sm StateManager) UpdateContainerStatus(status string) error {
	allowed := map[string]bool{
		"creating": true,
		"created":  true,
		"running":  true,
		"stopped":  true,
	}

	if !allowed[status] {
		return fmt.Errorf("unknown state: %s", status)
	}

	// current specs.State
	stateStruct, err := getContainerStateFileContentsAsStateStruct(sm.containerID)
	if err != nil {
		return err
	}

	// updated specs.State in memory
	stateStruct.Status = specs.ContainerState(status)
	updatedStateByteSlice, err := json.Marshal(stateStruct)
	if err != nil {
		return err
	}

	err = os.WriteFile(getContainerStateFileName(sm.containerID), updatedStateByteSlice, 0775)
	if err != nil {
		logrus.Errorf("could not update container state: %v", err)
		return err
	}

	return nil
}

func (sm StateManager) printStateFile() {
	contents, err := getContainerStateFileContents(sm.containerID)
	if err != nil {
		logrus.Errorf("unable to print container state contents")
		return
	}
	print(string(contents))
}

func CreateAndInitStateFile(containerID string, state *specs.State) error {
	dir := getContainerDirectoryName(containerID)
	stateFile := getContainerStateFileName(containerID)

	// create the stateFileLocation directory
	err := os.MkdirAll(dir, 0775)
	if err != nil {
		logrus.Fatalf("could not create directory %s: %v", dir, err)
		return err
	}

	_, err = os.Create(stateFile)
	if err != nil {
		logrus.Fatalf("could not create file %s: %v", stateFile, err)
		return err
	}

	stateJSONData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = os.WriteFile(stateFile, stateJSONData, 0666)
	if err != nil {
		logrus.Fatalf("could not write to state file %v: %v", state, err)
		return err
	}

	err = os.Chown(stateFile, 0, 0)
	if err != nil {
		logrus.Errorf("could not transfer ownership of %s to user root: %v", stateFile, err)
		return err
	}

	logrus.Infof("succesfull initialization of container state file at %s", stateFile)

	return nil
}

// UpdateProcessID updates the process ID in the state.json
// file to the desired newPidVal.
//
// This is not defaulted to the current PID for testing reasons.
func (sm StateManager) UpdateProcessID(newPidVal int) error {
	state, err := getContainerStateFileContentsAsStateStruct(sm.containerID)
	if err != nil {
		return err
	}
	state.Pid = newPidVal

	stateJSONData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = os.WriteFile(getContainerStateFileName(sm.containerID), stateJSONData, 0666)
	if err != nil {
		logrus.Fatalf("could not write to state file %v: %v", state, err)
		return err
	}

	return nil
}

// UpdateBundle updates the bundle path the state.json
// file to the desired bundlePath.
func (sm StateManager) UpdateBundle(bundlePath string) error {
	state, err := getContainerStateFileContentsAsStateStruct(sm.containerID)
	if err != nil {
		return err
	}
	state.Bundle = bundlePath

	stateJSONData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = os.WriteFile(getContainerStateFileName(sm.containerID), stateJSONData, 0666)
	if err != nil {
		logrus.Fatalf("could not write to state file %v: %v", state, err)
		return err
	}

	return nil
}

func (sm StateManager) FetchState() (*specs.State, error) {
	containerStateJson := fmt.Sprintf("%s/%s/state.json", MrunStateGlobalDirectory, sm.containerID)
	contents, err := os.ReadFile(containerStateJson)
	if err != nil {
		return nil, err
	}
	var state specs.State
	err = json.Unmarshal(contents, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func (sm StateManager) PrintStateFile() {
	containerStateJson := fmt.Sprintf("%s/%s/state.json", MrunStateGlobalDirectory, sm.containerID)
	contents, _ := os.ReadFile(containerStateJson)
	fmt.Printf(string(contents))
}
