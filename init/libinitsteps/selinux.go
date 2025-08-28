package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type setSELinuxLabelsLink struct {
	next Step
}

func (s *setSELinuxLabelsLink) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	//TODO implement me
	panic("implement me")
}
