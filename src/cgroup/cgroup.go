// package to interact with cgroupv2
// on systems running systemd as init

package cgroup

import (
	"errors"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/cgroup/controller"
	"os"
)

const (
	SysFsCgroup = "/sys/fs/cgroup"

	// DefaultCgroupMountTarget
	// default cgroup mount target is the user-1000 user slice
	DefaultCgroupMountTarget = SysFsCgroup + "/user.slice/user-1000.slice/user@1000.service/user.slice/"
)

type Cgroup struct {
	Name string
	// memoryController
	// nil if controller not enabled
	memoryController *controller.MemoryController

	// pid
	pid *controller.PidController

	cgroup *controller.CgroupController
}

// Create
//
// Make an instance of the control group.
func (cg *Cgroup) Create() error {
	err := cg.CreateCgroupDir()
	if err != nil {
		return err
	}
	err = cg.UpdateSubControllers()
	if err != nil {
		return err
	}
	err = cg.UpdateSubControllers()
	if err != nil {
		return err
	}
	return nil
}

// UpdateSubControllers
//
// Write values to sub-controller files for all controllers
func (cg *Cgroup) UpdateSubControllers() error {

	return nil
}

// WriteToCgroupController
//
// Write a value to a controller in the cgroup by controller name.
func (cg *Cgroup) WriteToCgroupController(subControllerName string, writeValue string) error {
	controllerAbsPath := cg.GetCgroupAbsolutePath() + subControllerName

	subControllerFile, err := os.OpenFile(controllerAbsPath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Could not open subControllerFile file '%s': %v\n", controllerAbsPath, err)
	}

	_, err = subControllerFile.Write([]byte(writeValue))
	if err != nil {
		return fmt.Errorf("Could not write to subControllerFile file '%s': %v\n", controllerAbsPath, err)
	}

	return nil
}

// GetCgroupAbsolutePath
//
// Get the absolute path to the cgroup
func (cg *Cgroup) GetCgroupAbsolutePath() string {
	absPath := DefaultCgroupMountTarget + cg.Name
	return absPath
}

// CreateCgroupDir
//
// an instance of your cgroup at the desired file location
func (cg *Cgroup) CreateCgroupDir() error {
	// check if program is being run as root
	MustBeRoot()

	cgroupAbsPath := cg.GetCgroupAbsolutePath() // absolute path of the cgroup being created

	err := os.Mkdir(cgroupAbsPath, 0755)
	if err != nil {
		return fmt.Errorf("Could not create a cgroup at %s: %w\n", cgroupAbsPath, err)
	}

	return nil
}

// Destroy
//
// Remove a control group by cgroupName
func (cg *Cgroup) Destroy() error {
	MustBeRoot()

	cgroupAbsPath := cg.GetCgroupAbsolutePath()
	err := os.RemoveAll(cgroupAbsPath)
	if err != nil {
		return fmt.Errorf("Could not remove cgroup at %s \n", cg.Name)
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
