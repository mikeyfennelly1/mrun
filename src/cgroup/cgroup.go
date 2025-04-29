package cgroup

import (
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// InitCgroup creates a new control group for the container.
func InitCgroup(sliceName string, specResources specs.LinuxResources) (*cgroup2.Manager, error) {
	// get cgroup2.Resources obj from specs.LinuxResources obj
	resources := cgroup2.ToResources(&specResources)

	// create the control group as direct descendant of root user slice.
	m, err := cgroup2.NewSystemd("/", sliceName, -1, resources)
	if err != nil {
		return nil, err
	}

	err = m.Update(resources)
	if err != nil {
		return nil, err
	}

	return m, nil
}
