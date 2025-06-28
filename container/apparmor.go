package container

import "github.com/opencontainers/runtime-spec/specs-go"

type SetAppArmorLink struct {
	next ChainLink
}

func (s SetAppArmorLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (s SetAppArmorLink) SetNext(link ChainLink) {
	//TODO implement me
	panic("implement me")
}
