package libinit

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
)

type createFileSystemLink struct {
	next ExecutableInitStep
}

func (c createFileSystemLink) Execute(spec *specs.Spec) error {
	return nil
}

func (c createFileSystemLink) SetNext(item ExecutableInitStep) {
	c.next = item
}

func createFileSystem(spec specs.Spec) error {
	if os.Geteuid() != 0 {
		fmt.Printf("not superuser\n")
	}

	// iterate through spec.Mounts, mounting each visited item
	for index, mount := range spec.Mounts {
		err := mountInContainer(mount)
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

// 'masks' the mount point of path by mounting an empty & read only tmpfs
// filesystem on that mount point
func maskPath(path string) error {
	err := unix.Mount("tmpfs", path, "tmpfs", unix.MS_RDONLY, "size=0")
	if err != nil {
		return err
	}

	return nil
}

func mountInContainer(fileSystemToMount specs.Mount) error {
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

// gets a bitmask for mount options for a mount.
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

func maskPaths(paths []string) {
	for _, path := range paths {
		err := unix.Mount("tmpfs", path, "tmpfs", 0, "")
		if err != nil {
			logrus.Warn(err)
		}
	}
}

func readOnlyPaths(paths []string) {
	for _, path := range paths {
		err := os.Chmod(path, 0555)
		if err != nil {
			logrus.Warn(err)
		}
	}
}
