package src

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseConfig(t *testing.T) {
	spec, err := parseConfig("./config.json")
	require.NoError(t, err)
	fmt.Printf("%v", spec)
}
