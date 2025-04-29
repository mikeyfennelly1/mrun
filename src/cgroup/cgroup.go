package cgroup

import (
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// InitCgroup creates a new control group for the container
// and adds the containers init process to the cgroup.
func InitCgroup(hostname string, specResources specs.LinuxResources) error {
	// get cgroup2.Resources obj from specs.LinuxResources obj
	resources := cgroup2.ToResources(&specResources)

	// create the control group as direct descendant of root user slice.
	_, err := cgroup2.NewSystemd("/", hostname, -1, resources)
	if err != nil {
		return err
	}

	return nil
}
