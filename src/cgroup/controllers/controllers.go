// controllers.go
//
// Interact with controllers in a cgroup
//
// @author Mikey Fennelly

package controllers

type Controllers struct {
	// memory
	// nil if controller not initialized
	memory *memController
}
