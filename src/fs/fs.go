package fs

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
	"os"
)

func CreateFileSystem(spec specs.Spec) error {
	if os.Geteuid() != 0 {
		fmt.Printf("not superuser\n")
	}

	// mount the god damn mounts
	for index, mount := range spec.Mounts {
		err := mountInContainerFS(mount)
		if err != nil {
			// should there be an error,
			// iterate back through all mounted filesystems and unmount each
			for i := index; i >= 0; i-- {
				unix.Unmount(spec.Mounts[i].Destination, 0)
			}
			return err
		}
	}

	// mask the paths
	for _, path := range spec.Linux.MaskedPaths {
		err := maskPath(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func maskPath(path string) error {
	err := unix.Mount("tmpfs", path, "tmpfs", unix.MS_RDONLY, "size=0")
	if err != nil {
		return err
	}

	return nil
}

func mountInContainerFS(fileSystemToMount specs.Mount) error {
	// check if the specified mount point exists
	_, err := os.Stat(fileSystemToMount.Destination)
	if os.IsNotExist(err) {
		// if not, create the mount point
		err = os.MkdirAll(fileSystemToMount.Destination, 0775)
		if err != nil {
			return err
		}
	}

	// mount the filesystem to mount point (destination)
	err = unix.Mount(fileSystemToMount.Source,
		fileSystemToMount.Destination,
		fileSystemToMount.Type,
		getBitMaskForMountOptions(fileSystemToMount.Options),
		"")

	if err != nil {
		return err
	}

	return nil
}

func getBitMaskForMountOptions(mountOptions []string) uintptr {
	var result uintptr

	for _, option := range mountOptions {
		switch option {
		case "nosuid":
			result |= unix.MS_NOSUID
		case "noexec":
			result |= unix.MS_NOEXEC
		case "nodev":
			result |= unix.MS_NODEV
		case "synchronous":
			result |= unix.MS_SYNCHRONOUS
		case "remount":
			result |= unix.MS_REMOUNT
		case "bind":
			result |= unix.MS_BIND
		case "move":
			result |= unix.MS_MOVE
		case "private":
			result |= unix.MS_PRIVATE
		case "shared":
			result |= unix.MS_SHARED
		}
	}
	return result
}
