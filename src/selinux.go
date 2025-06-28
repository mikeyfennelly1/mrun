package src

import "github.com/opencontainers/runtime-spec/specs-go"

type SetSELinuxLabelsLink struct {
	next ChainLink
}

func (s SetSELinuxLabelsLink) Execute(spec *specs.Spec) {
	//TODO implement me
	panic("implement me")
}

func (s SetSELinuxLabelsLink) SetNext(item ChainLink) {
	//TODO implement me
	panic("implement me")
}
