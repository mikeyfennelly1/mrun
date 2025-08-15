package state

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_initializeStateSuccess(t *testing.T) {
	// ensure no error occurred
	_, err := NewContainerState("test")
	require.NoError(t, err)

	require.True(t, validStateFile())

	// cleanup resources
	err = os.RemoveAll("/var/run/mrun/test")
	if err != nil {
		panic("unable to remove directory /var/run/mrun/test")
	}
}

func Test_initializeStateReturnsErr(t *testing.T) {
	_, err := NewContainerState("")
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
// 1. creating
// 2. created
// 3. running
// 4. stopped
func Test_updateStateValidParams(t *testing.T) {
	// creating
	sm, err := NewContainerState("test")
	require.NoError(t, err)
	err = sm.UpdateContainerStatus("creating")
	require.NoError(t, err)
	sm.printStateFile()
	err = sm.UpdateContainerStatus("running")
	require.NoError(t, err)
	sm.printStateFile()
	err = sm.UpdateContainerStatus("stopped")
	require.NoError(t, err)
	sm.printStateFile()
	err = DeleteState("test")
	require.NoError(t, err)
}

func Test_updateProcessId(t *testing.T) {
	sm, err := NewContainerState("test")
	require.NoError(t, err)
	sm.printStateFile()
	err = sm.UpdateProcessID(22)
	require.NoError(t, err)
	sm.printStateFile()
	err = DeleteState("test")
	require.NoError(t, err)
}

func Test_updateBundle(t *testing.T) {
	sm, err := NewContainerState("test")
	require.NoError(t, err)
	sm.printStateFile()
	err = sm.UpdateBundle("/test/bundle/path")
	require.NoError(t, err)
	sm.printStateFile()
	err = DeleteState("test")
	require.NoError(t, err)
}
