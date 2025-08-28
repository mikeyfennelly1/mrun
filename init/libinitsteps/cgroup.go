package libinitsteps

import (
	"fmt"
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	MrunCgroupSlice = "/"
)

// initCgroupStep is the cgroup initializing implementation of Step.
// It is the only interface for initializing a cgroup.
type initCgroupStep struct{}

func (i initCgroupStep) Execute(spec *specs.Spec, stateManager *state.StateManager) error {
	_, err := createNewCgroupForContainer(stateManager.GetContainerID(), spec, stateManager)
	if err != nil {
		return err
	}

	err = moveCurrentPidToCgroup(stateManager.GetContainerID())
	if err != nil {
		return err
	}

	return nil
}

func createNewCgroupForContainer(containerID string, spec *specs.Spec, sm *state.StateManager) (*cgroup2.Manager, error) {
	// get cgroup2.Resources obj from specs.LinuxResources obj
	resources := cgroup2.ToResources(spec.Linux.Resources)

	res := cgroup2.Resources{}

	groupName := fmt.Sprintf("%s.slice", sm.GetContainerID())

	// create the control group as direct descendant of root user slice.
	m, err := cgroup2.NewSystemd("/", groupName, -1, &res)
	if err != nil {
		logrus.Errorf("NewSystemd slice creation during cgroup initialization failed with error: %v", err)
		logrus.Errorf("unable to create new cgroup at location: /sys/fs/cgroup/%s", getGroupNameForContainerID(containerID))
		return nil, err
	}
	sm.SetCgroupInitialized()
	logrus.Infof("new cgroup created at path: /sys/fs/cgroup/%s", getGroupNameForContainerID(containerID))

	// update the cgroup controllers to match the contents of the cgroup2.Resources object
	err = m.Update(resources)
	if err != nil {
		logrus.Errorf("unable to update cgroup")
		return nil, fmt.Errorf("failed to update cgroup for container: %v", err)
	}
	logrus.Infof("updated resources for cgroup: /sys/fs/cgroup/%s", getGroupNameForContainerID(containerID))

	return m, nil
}

// moveCurrentPidToCgroup moves this program's process into a control group.
func moveCurrentPidToCgroup(containerID string) error {
	pid := os.Getpid()

	m, err := cgroup2.LoadSystemd(MrunCgroupSlice, getGroupNameForContainerID(containerID))
	if err != nil {
		logrus.Errorf("unable to load cgroup")
		return err
	}
	logrus.Infof("loaded cgroup: /sys/fs/cgroup/%s", getGroupNameForContainerID(containerID))

	err = m.AddProc(uint64(pid))
	if err != nil {
		logrus.Errorf("unable to add current process to cgroup")
		return err
	}
	logrus.Infof("updated resources for cgroup: /sys/fs/cgroup/%s", getGroupNameForContainerID(containerID))

	return nil
}

func getGroupNameForContainerID(containerID string) string {
	return fmt.Sprintf("%s.slice", containerID)
}

func deleteCgroup(containerID string) error {
	m, err := cgroup2.LoadSystemd(MrunCgroupSlice, getGroupNameForContainerID(containerID))
	if err != nil {
		return err
	}
	err = m.Delete()
	if err != nil {
		return err
	}

	return nil
}
