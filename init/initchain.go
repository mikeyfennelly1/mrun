package init

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

type ExecutableInitStep interface {
	Execute(spec *specs.Spec) error
}
