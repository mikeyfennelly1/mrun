package src

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateStateFile(t *testing.T) {
	m := GetContainerManager("rand")
	err := m.CreateAndInitStateFile()
	require.NoError(t, err)
}
