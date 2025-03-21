package controller

type ControllerProfile struct {
	// memory
	// nil if controller not enabled
	memory *memController

	// pid
	pid *pidController

	cgroup *cgroupController
}

func (cp *ControllerProfile) GetEnabledControllers() []ControllerId {
	var enabledControllers []ControllerId

	if cp.memory != nil {
		enabledControllers = append(enabledControllers, memoryControllerID)
	}
	if cp.pid != nil {
		enabledControllers = append(enabledControllers, pidControllerID)
	}
	if cp.cgroup != nil {
		enabledControllers = append(enabledControllers, cgroupControllerID)
	}

	return enabledControllers
}
