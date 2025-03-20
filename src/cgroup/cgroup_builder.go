//
// cgroup_builder.go
//
// An API to create a control group based on the systemd
// /sys/fs/cgroup structure
//
// @author Mikey Fennelly

package cgroup

import (
	"fmt"
	"github.com/moby/sys/mount"
	"syscall"
)

type CgroupType int

const (
	// Cgroup Types
	Domain CgroupType = iota
	Threaded
	DefaultMountPoint string = SysFsCgroup + "/user.slice/user-1000.slice/user@1000.service/user.slice/"
)

// Cgroup
// Represents a cgroup filesystem.
type Cgroup struct {
	// name of the control group that you want to create
	name string

	// cgroupType
	// the type of cgroup. can be either domain/threaded
	// https://docs.kernel.org/admin-guide/cgroup-v2.html#threads
	cgroupType CgroupType

	// mountPoint
	// mount point of the cgroup filesystem to create
	// https://docs.kernel.org/admin-guide/cgroup-v2.html#basic-operations
	//
	// in a system with systemd this is typically a '.slice' subdirectory
	/////////
	// providing no mountPoint options for early versions
	/////////

	// controllers
	// the controllers that you want to have in the control group
	// https://docs.kernel.org/admin-guide/cgroup-v2.html#controllers
	controllers *Controllers

	// pids
	// the processes that you want to be in the cgroup
	pids []int
}

// create
// Make an instance of the cgroup that you want
func (cg *Cgroup) create() error {
	return nil
}

// getAbsolutePath
// get the absolute path to the cgroup
func (cg *Cgroup) getAbsolutePath(uid int) string {
	absolutePath := DefaultMountPoint + cg.name
	return absolutePath
}

func (cg *Cgroup) mount() error {
	mountType := "cgroup2"  // type of filesystem to mount
	target := "/" + cg.name // target mount filesystem

	err := mount.Mount(DefaultMountPoint, target, mountType, "")
	if err != nil {
		return fmt.Errorf("Could not mount cgroup filesystem: %s\n", err.Error())
	}

	return nil
}

// movePidsToCg
// move processes (cg.pids) into a control group
func (cg *Cgroup) movePidsToCg() error {
	return nil
}
