package libstate

import (
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/containerd/cgroups/v3/cgroup2/stats"
	"github.com/mikeyfennelly1/mrun/libinit"
)

func GetResourceUsageInformation(containerID string) (*stats.Metrics, error) {
	cg, err := cgroup2.LoadSystemd(libinit.MrunCgroupSlice, containerID)
	if err != nil {
		return nil, err
	}

	stat, err := cg.Stat()
	if err != nil {
		return nil, err
	}

	return stat, nil
}
