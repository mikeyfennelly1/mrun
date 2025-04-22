package cgroup

import (
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func ConfigureCgroups(spec specs.LinuxResources) error {
	systemd, err := cgroup2.IOType()
	if err != nil {
		return
	}
}
