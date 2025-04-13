package main

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/namespace"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"syscall"
)

var rootfsPath string

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
		err := namespace.RestartInNewNS("chroot")
		if err != nil {
			return
		}
	},
}

func startSh() {
	cmd := exec.Command("/bin/sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return
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

		err = syscall.Chroot("./rootfs")
		if err != nil {
			return
		}
		err = os.Chdir("./rootfs")
		if err != nil {
			fmt.Printf("error changing directory to rootfs: %v", err)
			return
		}

		startSh()
	},
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	rootCmd.AddCommand(startCommand)
	rootCmd.AddCommand(chrootCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
