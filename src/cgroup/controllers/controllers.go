// controllers.go
//
// Interact with controllers in a cgroup
//
// @author Mikey Fennelly

package controllers

type Controller interface {
	// write the value of the controller to a file
	writeControllerValueToTarget(cgroupFsRoot string, controllerFilename string) error
}

type ControllerProfile struct {
	// memory
	// nil if controller not initialized
	memory *memController

	// pid
	pid *pidController

	cgroup *cgroupController
}
