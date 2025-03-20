package cgroup

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const SysFsCgroup = "/sys/fs/cgroup"

var mutex sync.Mutex

// PrintSysFsCgroupTree
//
// Print the directory structure of /sys/fs/cgroup
func PrintSysFsCgroupTree(level int) error {
	return filepath.WalkDir(SysFsCgroup, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Indentation for directory depth
		fmt.Printf("%s%s\n", string(' ', level*2), d.Name())
		if d.IsDir() {
			level++
		}
		return nil
	})
}
