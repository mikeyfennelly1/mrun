package libstate

import (
	"sync"
)

var lock = &sync.Mutex{}

type ContainerMetadata struct {
	containerName  string
	containerId    string
	bundleLocation string
}

var containerMetadata *ContainerMetadata

// getContainerMetadata gets the singleton ContainerMetadata
// object and returns to the calling client.
func getContainerMetadata() *ContainerMetadata {
	if containerMetadata == nil {
		lock.Lock()
		defer lock.Unlock()
		if containerMetadata == nil {
			containerMetadata = &ContainerMetadata{}
		}
	}

	return containerMetadata
}
