package test

import (
	"github.com/mikeyfennelly1/mrun/src/fs"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMountCgroup
// Test Mounting a cgroup filesystem at a specified
func TestMountCgroup(t *testing.T) {
	err := mount.MountCgroup("testie")
	assert.NoError(t, err)
}
