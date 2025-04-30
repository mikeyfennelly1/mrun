package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/cgroup"
	"github.com/mikeyfennelly1/mrun/src/fs"
	"github.com/mikeyfennelly1/mrun/src/proc"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"syscall"
)

var Chroot = &cobra.Command{
	Use:   "chroot",
	Short: "change the root for the binary.",
	Run: func(cmd *cobra.Command, args []string) {
		var spec specs.Spec
		jsonContent, err := os.ReadFile("./config.json")
		if err != nil {
			fmt.Printf("error reading config: %v", err)
			return
		}
		err = json.Unmarshal(jsonContent, &spec)
		if err != nil {
			fmt.Printf("error creating unmarshalling JSON: %v", err)
			return
		}

		m, err := cgroup.InitCgroup("test-container.slice", *spec.Linux.Resources)
		if err != nil {
			logrus.Errorf("could not initialize cgroup for container: %v\n", err)
			return
		}

		err = m.AddProc(uint64(os.Getpid()))
		if err != nil {
			logrus.Errorf("could not add this process to cgroup for container: %v\n", err)
			return
		}

		err = syscall.Chroot("./rootfs")
		if err != nil {
			return
		}
		err = os.Chdir("./rootfs")
		if err != nil {
			fmt.Printf("error changing directory to rootfs: %v", err)
			return
		}

		uid := int(spec.Process.User.UID)
		gid := int(spec.Process.User.GID)
		err = syscall.Setuid(uid)
		if err != nil {
			logrus.Warn("unable to set uid of process in container to %d\n", uid)
		}
		err = syscall.Setgid(gid)
		if err != nil {
			logrus.Warn("unable to set gid of process in container to %d\n", uid)
		}

		proc.SetRLIMITsForProcess(spec.Process.Rlimits)

		proc.SetEnvVars(spec.Process.Env)

		err = syscall.Sethostname([]byte(spec.Hostname))
		if err != nil {
			logrus.Warn(err)
		}

		err = fs.CreateFileSystem(spec)
		execSh()
	},
}

func execSh() {
	shell := "/bin/sh"
	args := []string{shell}
	env := syscall.Environ()

	err := syscall.Exec(shell, args, env)
	if err != nil {
		fmt.Printf("error execing shell: %v\n", err)
	}
}
