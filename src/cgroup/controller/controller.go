// Definitions of what a controller is and it's interfaces

package controller

type ControllerId int

// Controller
// Interface for a controller
type Controller interface {
	// GetSubControllerUpdates
	//
	// Get a map of controller filenames and string values
	// of what to write to those controller files to instantiate
	// that controller.
	GetSubControllerUpdates() []map[controllerFilename]string
}
