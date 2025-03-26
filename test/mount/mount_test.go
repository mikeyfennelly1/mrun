package test

import (
	"github.com/mikeyfennelly1/mrun/src/mount"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMountCgroup(t *testing.T) {
	err := mount.MountCgroup("testie")
	assert.NoError(t, err)
}
