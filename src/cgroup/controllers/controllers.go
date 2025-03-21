// controllers.go
//
// Interact with controllers in a cgroup
//
// @author Mikey Fennelly

package controllers

type ControllerId int

const (
	memoryControllerID ControllerId = iota
	pidControllerID
	cgroupControllerID
)

type Controller interface {
	enable() error
	setControllerValues() error
}

type ControllerProfile struct {
	// memory
	// nil if controller not enabled
	memory *memController

	// pid
	pid *pidController

	cgroup *cgroupController
}

func (cp *ControllerProfile) getEnabledControllers() []ControllerId {
	var enabledControllers []ControllerId

	if cp.memory != nil {
		enabledControllers = append(enabledControllers, memoryControllerID)
	}
	if cp.pid != nil {
		enabledControllers = append(enabledControllers, pidControllerID)
	}
	if cp.cgroup != nil {
		enabledControllers = append(enabledControllers, cgroupControllerID)
	}

	return enabledControllers
}
