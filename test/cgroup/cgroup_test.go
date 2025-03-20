package test

import (
	"github.com/mikeyfennelly1/mrun/src/cgroup"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCgroupWithoutPermissions(t *testing.T) {
	result := cgroup.CreateUserSliceCgroup(cgroup.THREADED, "test-cgroup")

	require.Error(t, result)
}

func TestCreateCgroupWithPermissions(t *testing.T) {
	err := cgroup.CreateUserSliceCgroup(cgroup.THREADED, "test-cgroup")
	require.NoError(t, err)
}
