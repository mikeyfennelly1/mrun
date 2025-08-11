package libinitsteps

import (
	"github.com/mikeyfennelly1/mrun/init"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"syscall"
)

type chrootLink struct {
	next init.Step
}

func (c chrootLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c chrootLink) SetNext(item init.Step) {
	//TODO implement me
	c.next = item
}

type changeProcessDirToNewRootLink struct {
	next init.Step
}

func (c changeProcessDirToNewRootLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c changeProcessDirToNewRootLink) SetNext(item init.Step) {
	//TODO implement me
	panic("implement me")
}

type setUsersAndGroupsLink struct {
	next init.Step
}

func (s setUsersAndGroupsLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (s setUsersAndGroupsLink) SetNext(item init.Step) {
	//TODO implement me
	s.next = item
}

type setRLIMITLink struct {
	next init.Step
}

func (s setRLIMITLink) Execute(spec *specs.Spec) error {
	setRLIMITsForProcess(spec.Process.Rlimits)
	return nil
}

func (s setRLIMITLink) SetNext(item init.Step) {
	//TODO implement me
	s.next = item
}

type setHostnameLink struct {
	next init.Step
}

func (s setHostnameLink) Execute(spec *specs.Spec) error {
	err := syscall.Sethostname([]byte(spec.Hostname))
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

func (s setHostnameLink) SetNext(item init.Step) {
	//TODO implement me
	panic("implement me")
}

type execBinaryLink struct {
	next init.Step
}

func (e execBinaryLink) Execute(spec *specs.Spec) error {
	//TODO implement me.
	// This should also execute shell as a default. Research if that is OCI compliant.
	panic("implement me")
	return nil
}

func (e execBinaryLink) SetNext(item init.Step) {
	//TODO implement me
	panic("implement me")
}
