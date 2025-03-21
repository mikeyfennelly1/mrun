// controllers.go
//
// Interact with controllers in a cgroup
//
// @author Mikey Fennelly

package controllers

type Controller interface {
	enable() error
	setControllerValues() error
}

type ControllerProfile struct {
	// memory
	// nil if controller not initialized
	memory *memController

	// pid
	pid *pidController

	cgroup *cgroupController
}
