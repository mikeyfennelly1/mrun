// package to interact with cgroupv2
// on systems running systemd as init

package cgroup

import (
	"errors"
	"fmt"
	"os"
)

const (
	SysFsCgroup = "/sys/fs/cgroup"

	// default cgroup mount target is the
	// user-1000 user slice
	DefaultCgroupMountTarget = SysFsCgroup + "/user.slice/user-1000.slice/user@1000.service/user.slice/"
)

// CreateCgroup
//
// Create a cgroup from OCI spec JSON
func CreateCgroup(cgroupName string) error {
	// check if program is being run as root
	MustBeRoot()

	cgroupAbsPath := DefaultCgroupMountTarget + cgroupName // absolute path of the cgroup being created

	err := os.Mkdir(cgroupAbsPath, 0755)
	if err != nil {
		return fmt.Errorf("Could not create a cgroup at %s: %w\n", cgroupAbsPath, err)
	}

	return nil
}

// DestroyCgroup
//
// Remove a control group by cgroupName
func DestroyCgroup(cgroupName string) error {
	MustBeRoot()

	cgroupAbsPath := DefaultCgroupMountTarget + cgroupName
	err := os.RemoveAll(cgroupAbsPath)
	if err != nil {
		return fmt.Errorf("Could not remove cgroup at %s \n", cgroupName)
	}

	return nil
}

// must
//
// Checker function to check if a value is an error.
//
// In the case that the value is an error, must() causes
// a program panic (non-0 exit)
func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// ErrNotRoot
//
// Error to signify that a user is not the root user.
var ErrNotRoot = errors.New("You must be root.\n")

// MustBeRoot
//
// Check that a user has uid of root.
//
//	return error if not.
func MustBeRoot() {
	if os.Geteuid() != 0 {
		panic(ErrNotRoot)
	}

	return
}
