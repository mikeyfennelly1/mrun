//go:generate mockgen -source=step.go -destination=../mocks/step.go -package=mocks
package libinitsteps

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Step is used to represent an executable step in an init
// chain/sequence of steps.
// Given a spec, it does it's duties, and reports an error should
// one happen.
type Step interface {
	Execute(spec *specs.Spec) error
}

func StepFactory[T any](specSubObj T) (Step, error) {
	switch any(specSubObj).(type) {
	case specs.Root:
		return &chrootStep{}, nil
	case string:
		return &createFileSystemStep{}, nil
	case string:
		return &execBinaryStep{}, nil
	case string:
		return &InitCgroupStep{}, nil
	case string:
		return &namespaceStep{}, nil
	case string:
		return &setAppArmorStep{}, nil
	case string:
		return &setEnvVarsStep{}, nil
	case string:
		return &setRLIMITStep{}, nil
	case string:
		return &setSELinuxLabelsLink{}, nil
	case string:
		return &setUsersAndGroupsStep{}, nil
	default:
		return nil, fmt.Errorf("unknown type for object: %v", specSubObj)
	}
}
