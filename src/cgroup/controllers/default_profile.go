package controllers

var DefaultControllerProfile = ControllerProfile{
	memory: &DefaultMemController,
	pid:    &DefaultPidController,
	cgroup: &DefaultCgroupController,
}
