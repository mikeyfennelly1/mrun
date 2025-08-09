package libinit

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type ChainLink interface {
	Execute(spec *specs.Spec) error
	SetNext(next ChainLink)
}
