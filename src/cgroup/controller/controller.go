// Definitions of what a controller is and it's interfaces

package controller

type ControllerId int

// Controller
// Interface for a controller
type Controller interface {
	// write the controller name to the cgroup.controllers file of the cgroup
	// this enables all controller subtypes to be enabled on that cgroup
	enable() error

	// write the value for each subcontroller to that subcontroller file
	setSubControllerValues() error

	// getControllerId
	// Get the controllerID of the controller
	getControllerId() ControllerId
}

const (
	memoryControllerID ControllerId = iota
	pidControllerID
	cgroupControllerID
)
