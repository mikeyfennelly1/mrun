package src

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type ChainLink interface {
	Execute(spec *specs.Spec)
	SetNext(ChainLink)
}
