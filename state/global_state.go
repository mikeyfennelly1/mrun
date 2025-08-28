//go:generate mockgen -source=subsystem.go -destination=../mocks/subsystem.go -package=mocks

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
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"os"
	"path/filepath"
)

const (
	MrunStateGlobalDirectory = "/var/run/mrun/"
	ociVersion               = "0.2.0"
)

var stateSingleton *specs.State = nil

// NewContainerState creates a directory in state
// directory /var/run/mrun/<container-id> and a file
// state.json in that directory. i.e
//
// /var/run/mrun/<container-id>/state.json
func NewContainerState(containerID string) (*StateManager, error) {
	if containerID == "" {
		return nil, fmt.Errorf("containerId is invalid")
	}
	containerDirname := fmt.Sprintf("%s%s", MrunStateGlobalDirectory, containerID)

	err := os.MkdirAll(containerDirname, 0775)
	if err != nil {
		logrus.Errorf("error while creating directory for managing state of container: %v", err)
		return nil, err
	}

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
	err = CreateAndInitStateFile(containerID, &state)
	if err != nil {
		return nil, err
	}

	return &StateManager{
		containerID: containerID,
	}, nil
}

func GetStateManager(containerID string) *StateManager {
	return &StateManager{containerID: containerID}
}

func StateFileExists(containerID string) bool {
	info, _ := os.Stat(fmt.Sprintf("%s/%s/state.json", MrunStateGlobalDirectory, containerID))
	if info != nil {
		return true
	}
	return false
}

func DeleteState(containerID string) error {
	err := os.RemoveAll(getContainerDirectoryName(containerID))
	if err != nil {
		return err
	}
	return nil
}

// validStateFile checks if the state file is valid according to the schema.
func validStateFile() bool {
	// Resolve the schema path to an absolute path
	absSchemaPath, err := filepath.Abs("../oci/state-schema.json")
	if err != nil {
		log.Fatalf("Failed to resolve schema path: %v", err)
	}

	// Resolve the document path to an absolute path
	absDocPath, err := filepath.Abs("/var/run/mrun/test/state.json")
	if err != nil {
		log.Fatalf("Failed to resolve document path: %v", err)
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + absSchemaPath)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + absDocPath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.Fatalf("Error during validation: %v", err)
	}

	return result.Valid()
}
