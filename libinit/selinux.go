package libinit

import "github.com/opencontainers/runtime-spec/specs-go"

type setSELinuxLabelsLink struct {
	next ExecutableInitStep
}

func (s setSELinuxLabelsLink) Execute(spec *specs.Spec) {
	//TODO implement me
	panic("implement me")
}

func (s setSELinuxLabelsLink) SetNext(item ExecutableInitStep) {
	//TODO implement me
	panic("implement me")
}
