package fs

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
)

/**
To create a container filesystem.

Need to have a root dir.
Mount a /proc in that directory.


*/

func CreateFileSystem(spec specs.Spec) error {
	if os.Geteuid() != 0 {
		fmt.Printf("not superuser\n")
	}
	fmt.Printf("EUID: %d\n", os.Geteuid())

	// mount the god damn mounts
	for index, mount := range spec.Mounts {
		err := mountInContainerFS(spec.Root.Path, mount)
		if err != nil {
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

func mountInContainerFS(rootPath string, fileSystemToMount specs.Mount) error {
	info, err := os.Stat(fileSystemToMount.Destination)
	if info == nil {
		err = os.MkdirAll(fileSystemToMount.Destination, 0775)
		if err != nil {
			return err
		}
	}

	err = unix.Mount(fileSystemToMount.Source,
		fileSystemToMount.Destination,
		fileSystemToMount.Type,
		getBitMaskForMountOptions(fileSystemToMount.Options),
		"")

	if err != nil {
		if err.Error() == "device or resource busy" {
			cmd := exec.Command("lsof", "+D", fileSystemToMount.Destination)
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Printf("device or resource busy, lsof for file:")

		}
		return fmt.Errorf("%s: %v", fileSystemToMount.Destination, err)
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
