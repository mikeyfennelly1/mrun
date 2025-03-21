package controller

import "strconv"

// CgroupController
//
// Represents cgroup.<sub-controller> files.
type CgroupController struct {
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
	memory  bool
	pids    bool
	cpu     bool
	io      bool
	netCls  bool
	netPrio bool
	rdma    bool
	devices bool
}

// GetSubControllerUpdates
//
// Get the filenames, and corresponding updates
// for this controller's sub-controllers
//
// returns map of filenames to update values
func (cgc *CgroupController) GetSubControllerUpdates() []map[controllerFilename]string {
	var fileWriteValMapSlice []map[controllerFilename]string

	// check if cgroup.subtree_control is enabled
	//
	// if so write all values of the sutree_control controller
	// to cgroup.subtree control
	if cgc.subtreeControl != nil {
		for _, value := range cgc.subtreeControl.getSubtreeControlWriteVals() {
			fileWriteValMapSlice = append(fileWriteValMapSlice, map[controllerFilename]string{
				"cgroup.subtree_control": value,
			})
		}
	}

	// check if cgroup.freeze is true, if so write 1 to cgroup.freeze, else write 0
	if cgc.freeze {
		fileWriteValMapSlice = append(fileWriteValMapSlice, map[controllerFilename]string{
			"cgroup.freeze": "1",
		})
	} else {
		fileWriteValMapSlice = append(fileWriteValMapSlice, map[controllerFilename]string{
			"cgroup.freeze": "0",
		})
	}

	// check if process slice is empty, if not add contents as string to fileWriteValMapSlice
	if len(cgc.procs) != 0 {
		for _, pid := range cgc.procs {
			fileWriteValMapSlice = append(fileWriteValMapSlice, map[controllerFilename]string{
				"cgroup.procs": strconv.Itoa(pid),
			})
		}
	}

	return fileWriteValMapSlice
}

// getSubtreeControlWriteVals
//
// Get values that must be written to cgroup.subtree_control
// to update the cgroup.
func (stc *subtreeControl) getSubtreeControlWriteVals() []string {
	var finalVal []string
	if stc.memory {
		finalVal = append(finalVal, "memory")
	}
	if stc.pids {
		finalVal = append(finalVal, "pids")
	}
	if stc.cpu {
		finalVal = append(finalVal, "cpu")
	}
	if stc.io {
		finalVal = append(finalVal, "io")
	}
	if stc.netCls {
		finalVal = append(finalVal, "net_cls")
	}
	if stc.netPrio {
		finalVal = append(finalVal, "net_prio")
	}
	if stc.rdma {
		finalVal = append(finalVal, "rdma")
	}
	if stc.devices {
		finalVal = append(finalVal, "devices")
	}
	return finalVal
}
