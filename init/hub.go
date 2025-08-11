package init

import "github.com/mikeyfennelly1/mrun/init/libinit"

const configJsonPath = "./config.json"

type HubInterface interface {
	// ContainerIsCreating decides if this container is either
	// is in the process of being created already or not.
	//
	// This is necessary as after exec() (which is required to enter new namespaces)
	// the process has lost all of its memory. At this crucial point (entering new namespaces)
	// we actually execute the 'mrun' binary, with the 'start' argument.
	//
	// This essentially means that from mrun's perspective, it has to check
	// "Am I continuing on the work of a previous process that was in the middle of creating a
	// container, or am I starting the creation a new container?"
	//
	// Thus, we ask the mrun state subsystem, if the state of the container that
	// we are creating 'CREATING'. If so, we proceed based on the state.json
	// to continue to create the container by finding the delta between the state.json
	// and the config.json (the manifest essentially)
	ContainerIsCreating() bool

	// GetSteps creates the init chain based on if it is currently
	// being created already or not and returns the first link
	GetSteps(isCreating bool) []libinit.ExecutableInitStep
}

type hub struct {
	// initChain is a linked list of items to execute
	// in sequence in order to complete this stage of init
	InitChain libinit.ExecutableInitStep
}
