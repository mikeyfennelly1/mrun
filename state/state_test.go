package state

import (
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_initializeStateSuccess(t *testing.T) {
	// ensure no error occurred
	_, err := initializeState("test")
	require.NoError(t, err)

	require.True(t, validStateFile())

	// cleanup resources
	err = os.RemoveAll("/var/run/mrun/test")
	if err != nil {
		panic("unable to remove directory /var/run/mrun/test")
	}
}

func Test_initializeStateReturnsErr(t *testing.T) {
	_, err := initializeState("")
	require.Error(t, err)
}

// Test_updateStateBadParams requires that an error be thrown if a
// parameter is passed that intends to update the state.json file for
// a container to a state that is invalid
//
// All values that are not one of ["creating", "created", "running", "stopped"]
// is considered invalid.
func Test_updateStateBadParams(t *testing.T) {

}

// Test_updateStateValidParams tests updating states in a state.json
// file to all valid options. It goes in the order:
// 1. Creating
// 2. Created
// 3. Running
// 4. Stopped
func Test_updateStateValidParams(t *testing.T) {
	manager, err := initializeState("test")
	require.NoError(t, err)

}

func Test_updateProcessId(t *testing.T) {
	panic("implement me")
}

func Test_updateBundle(t *testing.T) {
	panic("implement me")
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
