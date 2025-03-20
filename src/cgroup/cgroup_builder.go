//
// cgroup_builder.go
//
// An API to create a control group based on the systemd
// /sys/fs/cgroup structure
//
// @author Mikey Fennelly

package cgroup

type CgroupType int

const (
	// Cgroup Types
	Domain CgroupType = iota
	Threaded
)

// Cgroup
//
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
	mountPoint string

	// controllers
	// the controllers that you want to have in the control group
	// https://docs.kernel.org/admin-guide/cgroup-v2.html#controllers
	controllers *Controllers

	// pids
	// the processes that you want to be in the cgroup
	pids []int
}

// getAbsolutePath
//
// get the absolute path to the cgroup
func (cg *Cgroup) getAbsolutePath(uid int) *string {
	return nil
}

// create
// Make an instance of the cgroup that you want
func (cg *Cgroup) create() error {
	return nil
}
