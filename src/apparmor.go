package src

import "github.com/opencontainers/runtime-spec/specs-go"

type SetAppArmorLink struct {
	next ChainLink
}

func (s SetAppArmorLink) Execute(spec *specs.Spec) {
	//TODO implement me
	panic("implement me")
}

func (s SetAppArmorLink) SetNext(link ChainLink) {
	//TODO implement me
	panic("implement me")
}
