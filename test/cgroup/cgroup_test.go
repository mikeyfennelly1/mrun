package test

import (
	"github.com/mikeyfennelly1/mrun/src/cgroup"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateCgroup(t *testing.T) {
	// rootless testing should fail
	if os.Geteuid() != 0 {
		rootlessResult := cgroup.CreateCgroup("test-cgroup3")
		require.Error(t, rootlessResult)
	} else {
		rootFulResult := cgroup.CreateCgroup("test-cgroup3")
		require.NoError(t, rootFulResult)
	}
}
