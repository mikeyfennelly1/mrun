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

// StepFactory gets a Step for a given part of the
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
		return &restartInNewNamespacesStep{}, nil
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
	case "usersgroups":
		return &setUsersAndGroupsStep{}, nil
	case "hostname":
		return &setHostnameStep{}, nil
	default:
		return nil, fmt.Errorf("unknown step name: %s", stepName)
	}
}
