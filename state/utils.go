package state

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

func getContainerStateFileName(containerID string) string {
	return fmt.Sprintf("%s/state.json", getContainerDirectoryName(containerID))
}

func getContainerDirectoryName(containerID string) string {
	return fmt.Sprintf("%s%s", MrunStateGlobalDirectory, containerID)
}

func getContainerStateFileContents(containerID string) ([]byte, error) {
	contents, err := os.ReadFile(getContainerStateFileName(containerID))
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func getContainerStateFileContentsAsStateStruct(containerID string) (*specs.State, error) {
	stateByteArray, err := getContainerStateFileContents(containerID)
	if err != nil {
		logrus.Errorf("could not update container state: %v", err)
		return nil, err
	}

	// unmarshal byte slice into an understandable specs.Spec structure
	var spec specs.State
	err = json.Unmarshal(stateByteArray, &spec)
	if err != nil {
		logrus.Errorf("unable to unmarshal container statefile: %v", err)
		return nil, err
	}

	return &spec, nil
}
