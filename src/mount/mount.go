package mount

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func MountCgroup(cgroupName string) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must be root to perform this operation\n")
	}
	target := "/mnt/test-cgroup"

	err := unix.Mount(cgroupName, target, "cgroup2", uintptr(0), "")
	if err != nil {
		return err
	}

	return nil
}
