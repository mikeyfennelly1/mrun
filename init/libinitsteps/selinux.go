package libinitsteps

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type setSELinuxLabelsLink struct {
	next Step
}

func (s *setSELinuxLabelsLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
}
