package main

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/cgroup"
	"github.com/mikeyfennelly1/mrun/src/fs"
	"github.com/mikeyfennelly1/mrun/src/namespace"
	"github.com/mikeyfennelly1/mrun/src/proc"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/syndtr/gocapability/capability"
	"os"
	"syscall"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mrun", // The name of the command
	Short: "A low-level container runtime.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
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

		proc.SetAndApplyCapsetToCurrentPid(capability.INHERITABLE, spec.Process.Capabilities.Inheritable)

		proc.SetAndApplyCapsetToCurrentPid(capability.PERMITTED, spec.Process.Capabilities.Permitted)

		proc.SetAndApplyCapsetToCurrentPid(capability.EFFECTIVE, spec.Process.Capabilities.Effective)

		proc.SetAndApplyCapsetToCurrentPid(capability.AMBIENT, spec.Process.Capabilities.Ambient)

		err = namespace.RestartInNewNS("chroot")
		if err != nil {
			logrus.Fatalf("unable to enter new namesapces")
			logrus.Exit(1)
		}
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

var chrootCommand = &cobra.Command{
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

func main() {
	rootCmd.AddCommand(startCommand)
	rootCmd.AddCommand(chrootCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
