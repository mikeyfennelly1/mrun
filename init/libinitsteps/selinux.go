package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/init"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type setSELinuxLabelsLink struct {
	next init.Step
}

func (s setSELinuxLabelsLink) Execute(spec *specs.Spec) {
	//TODO implement me
	panic("implement me")
}

func (s setSELinuxLabelsLink) SetNext(item init.Step) {
	//TODO implement me
	panic("implement me")
}
