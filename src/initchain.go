package src

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type ChainItem interface {
	Execute(spec *specs.Spec)
	SetNext(ChainItem)
}
