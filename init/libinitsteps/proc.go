package libinitsteps

import (
	"os"
	"syscall"

	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

// chrootStep executes the chroot syscall and moves the
// process into the new root.
type chrootStep struct{}

func (c *chrootStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	// change cwd to the new root
	logrus.Tracef("changing current working directory to: %s", spec.Root.Path)
	err := os.Chdir(spec.Root.Path)
	if err != nil {
		logrus.Errorf("error changing current working directory: %v", err)
		return err
	}
	logrus.Tracef("changed directory to: %s", spec.Root.Path)

	// perform chroot syscall
	logrus.Infof("attempting to change fs root to: %s", spec.Root.Path)
	err = syscall.Chroot(spec.Root.Path)
	if err != nil {
		logrus.Errorf("an error occurred in changing root directory to %s for this process: %v", spec.Root.Path, err)
		return err
	}

	return nil
}

type setHostnameStep struct{}

func (s *setHostnameStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	err := syscall.Sethostname([]byte(spec.Hostname))
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

type execBinaryStep struct{}

func (e *execBinaryStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	//TODO implement me.
	// This should also execute shell as a default. Research if that is OCI compliant.
	panic("execBinaryStep: implement me")
	return nil
}
