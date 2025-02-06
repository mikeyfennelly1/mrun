package test

import (
	container_init "github.com/mikeyfennelly1/mrun/v2/container-init"
	"testing"
)

func TestUnmarshalLayer(t *testing.T) {
	var expected = &container_init.Layer{
		MediaType: "application/vnd.docker.image.rootfs.diff.tar.gzip",
		Size:      5417,
		Digest:    "sha256:7bc940808c30e57ff4f6c47330cc7e9a7afbc22531ba4f94363fb614652ccb94",
	}

	actual, _ := container_init.UnMarshalLayer(`{
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 5417,
         "digest": "sha256:7bc940808c30e57ff4f6c47330cc7e9a7afbc22531ba4f94363fb614652ccb94"
      }`)

	if *expected != *actual {
		t.Errorf("Assertation failed: expected %v, got %v", *expected, *actual)
	}
}
