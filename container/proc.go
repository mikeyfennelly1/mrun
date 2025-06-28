package container

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"syscall"
)

type ChrootLink struct {
	next ChainLink
}

func (c ChrootLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c ChrootLink) SetNext(item ChainLink) {
	//TODO implement me
	c.next = item
}

type ChangeProcessDirToNewRootLink struct {
	next ChainLink
}

func (c ChangeProcessDirToNewRootLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c ChangeProcessDirToNewRootLink) SetNext(item ChainLink) {
	//TODO implement me
	panic("implement me")
}

type SetUsersAndGroupsLink struct {
	next ChainLink
}

func (s SetUsersAndGroupsLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (s SetUsersAndGroupsLink) SetNext(item ChainLink) {
	//TODO implement me
	s.next = item
}

type SetRLIMITLink struct {
	next ChainLink
}

func (s SetRLIMITLink) Execute(spec *specs.Spec) error {
	SetRLIMITsForProcess(spec.Process.Rlimits)
	return nil
}

func (s SetRLIMITLink) SetNext(item ChainLink) {
	//TODO implement me
	s.next = item
}

type SetHostnameLink struct {
	next ChainLink
}

func (s SetHostnameLink) Execute(spec *specs.Spec) error {
	err := syscall.Sethostname([]byte(spec.Hostname))
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

func (s SetHostnameLink) SetNext(item ChainLink) {
	//TODO implement me
	panic("implement me")
}

type ExecBinaryLink struct {
	next ChainLink
}

func (e ExecBinaryLink) Execute(spec *specs.Spec) error {
	//TODO implement me.
	// This should also execute shell as a default. Research if that is OCI compliant.
	panic("implement me")
	return nil
}

func (e ExecBinaryLink) SetNext(item ChainLink) {
	//TODO implement me
	panic("implement me")
}
