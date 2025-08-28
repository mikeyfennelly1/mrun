package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type setAppArmorStep struct{}

func (s *setAppArmorStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	//TODO implement me
	panic("implement me")
	return nil
}
