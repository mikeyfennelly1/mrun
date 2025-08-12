package init

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type Step interface {
	Execute(spec *specs.Spec) error
}
