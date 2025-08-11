package init

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Init() (HubInterface, error) {
	// check if the config.json exists in the pwd
	if !configJsonExists() {
		logrus.Fatal("./config.json does not exist")
		os.Exit(1)
	}
	// read contents of config.json
	contents, err := getConfigJsonContents()
	if err != nil {
		logrus.Fatal("./config.json does not exist")
		os.Exit(1)
	}
	// validate that it is structurally correct

	return nil, nil
}

// configJsonExists uses Unix stat syscall to check if file
// metadata is retrievable
func configJsonExists() bool {
	// fail if stat information is not found
	s, err := os.Stat(configJsonPath)
	// if there is an error, fail
	if err != nil {
		return false
	}

	if s != nil {
		return true
	}
	return false
}

// getConfigJsonContents returns the bytearray contents
// of $(pwd)/config.json.
//
// This should only really be used after configJsonExists
// It does not care if there are no contents in the file
func getConfigJsonContents() (*[]byte, error) {
	fileContents, _ := os.ReadFile(configJsonPath)
	if fileContents != nil {
		return nil, fmt.Errorf("failed to read %s", configJsonPath)
	}

	return &fileContents, nil
}

// configIsValid checks if a passed config.json file
// is valid according to the config schema in the
// OCI specification.
func configIsValid(config *string) bool {
	//TODO: implement me
	panic("configIsValid: implement me")
	return false
}
