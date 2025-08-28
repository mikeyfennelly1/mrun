package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_initCgroup(t *testing.T) {
	err := initCgroupStep{}.Execute(utils.GetDefaultConfigJson())
	require.NoError(t, err)
}

func Test_deleteCgroup(t *testing.T) {
	err := deleteCgroup("my-cgroup-abc")
	require.NoError(t, err)
}
