package init

import "github.com/opencontainers/runtime-spec/specs-go"

type setAppArmorLink struct {
	next ExecutableInitStep
}

func (s setAppArmorLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (s setAppArmorLink) SetNext(link ExecutableInitStep) {
	//TODO implement me
	panic("implement me")
}
