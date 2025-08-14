package libinitsteps

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

// chrootStep executes the chroot syscall and moves the
// process into the new root.
type chrootStep struct{}

func (c *chrootStep) Execute(spec *specs.Spec) error {
	// perform chroot syscall
	err := syscall.Chroot(spec.Root.Path)
	if err != nil {
		logrus.Errorf("an error occurred in changing root directory for this process: %v", err)
		return err
	}

	// change cwd to the new root
	logrus.Tracef("changing current working directory to: %s", spec.Root.Path)
	err = os.Chdir(spec.Root.Path)
	if err != nil {
		logrus.Errorf("error changing current working directory: %v", err)
		return err
	}
	logrus.Tracef("changed directory to: %s", spec.Root.Path)
	return nil
}

type setUsersAndGroupsStep struct{}

func (s *setUsersAndGroupsStep) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

type SetHostnameStep struct{}

func (s *SetHostnameStep) Execute(spec *specs.Spec) error {
	err := syscall.Sethostname([]byte(spec.Hostname))
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

type execBinaryStep struct{}

func (e *execBinaryStep) Execute(spec *specs.Spec) error {
	//TODO implement me.
	// This should also execute shell as a default. Research if that is OCI compliant.
	panic("implement me")
	return nil
}
