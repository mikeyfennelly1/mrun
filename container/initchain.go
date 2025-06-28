package container

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type ChainLink interface {
	Execute(spec *specs.Spec)
	SetNext(next ChainLink)
}
