package container_init

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func UnpackGzipFile(gzFilePath, dstFilePath string) (int64, error) {
	gzFile, err := os.Open(gzFilePath)
	if err != nil {
		return 0, fmt.Errorf("open file %q to unpack: %w", gzFilePath, err)
	}
	dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return 0, fmt.Errorf("create destination file %q to unpack: %w", dstFilePath, err)
	}
	defer dstFile.Close()

	ioReader, ioWriter := io.Pipe()
	defer ioReader.Close()

	go func() { // goroutine leak is possible here
		gzReader, _ := gzip.NewReader(gzFile)
		// it is important to close the writer or reading from the other end of the
		// pipe or io.copy() will never finish
		defer func() {
			gzFile.Close()
			gzReader.Close()
			ioWriter.Close()
		}()

		io.Copy(ioWriter, gzReader)
	}()

	written, err := io.Copy(dstFile, ioReader)
	if err != nil {
		return 0, err // goroutine leak is possible here
	}

	return written, nil
}
