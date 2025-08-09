package libinit

import (
	"fmt"
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
)

const (
	MrunCgroupSlice = "/"
)

// InitCgroupLink is the cgroup initializing implementation of ChainLink.
// It is the only interface for initializing a cgroup.
type InitCgroupLink struct {
	next ChainLink
}

func (i *InitCgroupLink) Execute(spec *specs.Spec) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (i *InitCgroupLink) SetNext(next ChainLink) {
	i.next = next
}

// InitCgroup creates a new 'blank' control group for the container.
func InitCgroup(containerID string, spec specs.Spec) error {
	m, err := createNewCgroupForContainer(containerID, *spec.Linux.Resources)
	if err != nil {
		return fmt.Errorf("could not initialize cgroup for containerID %s: %v", containerID, err)
	}

	// update the created cgroup with resources defined in config.json
	resources := cgroup2.ToResources(spec.Linux.Resources)
	err = m.Update(resources)
	if err != nil {
		return fmt.Errorf("failed to update cgroup for container: %v", err)
	}

	return nil
}

// MoveCurrentPidToCgroup moves this program's process into a control group.
func MoveCurrentPidToCgroup(containerID string) error {
	pid := os.Getpid()

	m, err := cgroup2.LoadSystemd(MrunCgroupSlice, getGroupNameForContainerID(containerID))
	if err != nil {
		return err
	}

	err = m.AddProc(uint64(pid))
	if err != nil {
		return err
	}

	return nil
}

func createNewCgroupForContainer(containerID string, specResources specs.LinuxResources) (*cgroup2.Manager, error) {
	// get cgroup2.Resources obj from specs.LinuxResources obj
	resources := cgroup2.ToResources(&specResources)

	groupName := getGroupNameForContainerID(containerID)
	// create the control group as direct descendant of root user slice.
	m, err := cgroup2.NewSystemd(MrunCgroupSlice, groupName, -1, resources)
	if err != nil {
		return nil, err
	}

	err = m.Update(resources)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func getGroupNameForContainerID(containerID string) string {
	return fmt.Sprintf("%s.slice", containerID)
}
