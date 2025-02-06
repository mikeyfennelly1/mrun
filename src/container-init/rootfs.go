package container_init

import (
	"encoding/json"
	"fmt"
)

// This file is created to build a container root filesystem with
// overlayfs from an image's manifest.json
//
// see https://github.com/opencontainers/image-spec/blob/main/schema/image-manifest-schema.jsonhttps://opencontainers.org/schema/image/descriptor/mediaType

// filesystem layer
type Layer struct {
	// Media type defines what type of image
	MediaType string `json:"mediaType"`
	// Size of the filesystem layer
	Size uint32 `json:"size"
` // Digest SHA
	Digest string `json:"digest"`
}

// mount a child filesystem (childFS) to a target parent filesystem (parentFS)
func (parentFS *Layer) mount(childFS Layer, mountPoint string) error {
	fmt.Printf("childFs: %v, , mountPointL %v\n", childFS, mountPoint)
	return nil
}

func overlay(lowerdir string, upperdir string, workdir string) error {
	return nil
}

// iterate through Layer slice and perform a non-arbitrary
// operation on each Layer item in slice
func IterateLayers(layers []Layer) {
	for i, l := range layers {
		nextLayer := layers[i+1]
		l.mount(nextLayer, "/mount/point")
	}
}

// Marshal a JSON string to Layer format
func UnMarshalLayer(layerJSONString string) (*Layer, error) {
	var layer Layer
	err := json.Unmarshal([]byte(layerJSONString), &layer)
	return &layer, err
}
