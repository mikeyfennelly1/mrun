//go:generate mockgen -source=step.go -destination=../mocks/step.go -package=mocks
package libinitsteps

import (
	"fmt"

	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Step is used to represent an executable step in an init chain/sequence of steps.
// Given a spec, it does it's duties, and reports an error - should one happen.
type Step interface {
	Execute(spec *specs.Spec, stateManager *state.StateManager) error
}

// StepFactory gets a Step for a given part of the
// init chain based on the title for that step.
//
// A developer can consult the source code of this
// factory function for the titles available.
func StepFactory(stepName string) (Step, error) {
	switch stepName {
	case "chroot":
		return &chrootStep{}, nil
	case "create-fs":
		return &createFileSystemStep{}, nil
	case "exec-bin":
		return &execBinaryStep{}, nil
	case "init-cgroup":
		return &initCgroupStep{}, nil
	case "namespace":
		return &startInNewNamespacesStep{}, nil
	case "apparmor":
		return &setAppArmorStep{}, nil
	case "set-env":
		return &setEnvVarsStep{}, nil
	case "rlimit":
		return &setRLIMITStep{}, nil
	case "capset":
		return &applyCapsetStep{}, nil
	case "selinux":
		return &setSELinuxLabelsLink{}, nil
	case "hostname":
		return &setHostnameStep{}, nil
	default:
		return nil, fmt.Errorf("unknown step name: %s", stepName)
	}
}
