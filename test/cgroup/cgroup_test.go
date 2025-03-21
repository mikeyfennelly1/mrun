package test

import (
	"github.com/mikeyfennelly1/mrun/src/cgroup"
	"github.com/stretchr/testify/require"
	"os"
	"syscall"
	"testing"
)

const (
	RootUid        = 0
	DefaultUid     = 1000
	TestCgroupName = "test-cgroup3"
)

var testCgroup = cgroup.Cgroup{
	Name: "test-cgroup",
}

// TestDestroyCgroup
//
// Test destroying a control group.
func TestDestroyCgroup(t *testing.T) {
	result := testCgroup.Destroy()
	require.NoError(t, result)
}

// TestRootfulCgroupOperations
//
// Test creating a control group with rootful permissions (legal).
func TestRootfulCgroupOperations(t *testing.T) {
	cgroup.MustBeRoot()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got %v", r)
		}
	}()

	defer testCgroup.Destroy()

	result := testCgroup.CreateCgroupDir()
	defer require.NoError(t, result)
}

// TestRootLessCgroupOperationFailure
//
// Test attempting to perform an operation on a container without root
// permissions (illegal, causes program panic).
func TestRootLessCgroupOperationFailure(t *testing.T) {
	if os.Geteuid() == RootUid {
		syscall.Seteuid(DefaultUid)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but got none\n")
		}
	}()

	defer testCgroup.Destroy()

	// this should panic
	testCgroup.CreateCgroupDir()
}
