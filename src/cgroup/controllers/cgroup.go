package controllers

// cgroupController
type cgroupController struct {
	freeze         bool
	kill           bool
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
