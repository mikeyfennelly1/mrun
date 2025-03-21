// package to interact with cgroupv2
// on systems running systemd as init

package cgroup

import (
	"fmt"
	"os"
)

const (
	CgroupFsType = "cgroup2"

	SysFsCgroup = "/sys/fs/cgroup"

	// default cgroup mount target is the
	// user-1000 user slice
	DefaultCgroupMountTarget = SysFsCgroup + "/user.slice/user-1000.slice/user@1000.service/user.slice/"
)

// CreateCgroup
//
// Create a cgroup from OCI spec JSON
func CreateCgroup(cgroupName string) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must be root user to create a cgroup\n")
	}
	cgroupAbsPath := DefaultCgroupMountTarget + cgroupName
	err := os.Mkdir(cgroupAbsPath, 0755)
	if err != nil {
		return fmt.Errorf("Could not create a cgroup at %s: %w\n", cgroupAbsPath, err)
	}
	return nil
}
