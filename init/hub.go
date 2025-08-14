//go:generate mockgen -source=hub.go -destination=../mocks/hub.go -package=mocks
package init

import (
	"github.com/mikeyfennelly1/mrun/init/libinitsteps"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const configJsonPath = "./config.json"

// GetSteps creates a slice of Step implementations
// that correspond to the steps required between now and the
// end of the init stage.
//
// The end of the init stage can be either:
//  1. Executing mrun start again, wherein the
//     mrun process can determine if it is already
//     in the process of initializing a container.	.
//  2. Executing the application binary.
//
// @param containerID: *string | null
//
//	If this is left null, it is assumed
//	that we are at the first stage of the init process. i.e.
//	that the container does not exist yet.
//	If it is not null, and a valid containerID has been
//	passed, GetSteps retrieves a StateManager object for the
//	statefile that corresponds with the containerID via the .
//
// @param
func GetSteps(containerID *string, configJson *specs.Spec) {

}

type hub struct {
	// initChain is a linked list of items to execute
	// in sequence in order to complete this stage of init
	steps []libinitsteps.Step
}
