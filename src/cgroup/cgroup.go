package cgroup

import (
	"fmt"
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func ConfigureCgroups(specResources specs.LinuxResources) error {
	fmt.Printf("%v\n", specResources)
	resources := cgroup2.ToResources(&specResources)
	fmt.Printf("%v\n", resources)
	return nil
}
