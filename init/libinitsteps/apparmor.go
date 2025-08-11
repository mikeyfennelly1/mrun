package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/init"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type setAppArmorLink struct {
	next init.Step
}

func (s setAppArmorLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (s setAppArmorLink) SetNext(link init.Step) {
	//TODO implement me
	panic("implement me")
}
