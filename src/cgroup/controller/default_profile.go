package controller

var DefaultControllerProfile = ControllerProfile{
	memory: &DefaultMemController,
	pid:    &DefaultPidController,
	cgroup: &DefaultCgroupController,
}
