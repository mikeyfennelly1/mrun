package controllers

// cgroupController
type cgroupController struct {
	freeze         bool
	maxDepth       int32
	maxDescendants int32
	//pressure
	procs          []int
	subtreeControl *subtreeControl
}

// subtreeControl
// enable/disable usage of controllers in subtree
// cgroups
type subtreeControl struct {
	memory   bool
	pids     bool
	cpu      bool
	io       bool
	net_cls  bool
	net_prio bool
	rdma     bool
	devices  bool
}

var DefaultCgroupController = cgroupController{
	freeze:         false,
	maxDepth:       3,
	maxDescendants: 3,
	procs:          nil,
	subtreeControl: &defaultSubtreeControl,
}

var defaultSubtreeControl = subtreeControl{
	memory:   true,
	pids:     true,
	cpu:      false,
	io:       false,
	net_cls:  false,
	net_prio: false,
	rdma:     false,
	devices:  false,
}
