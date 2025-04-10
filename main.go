package main

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/namespace"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mrun", // The name of the command
	Short: "A low-level container runtime.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// command for removing an eefenn-cli command
var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Start a container.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// command for removing an eefenn-cli command
var isoBashCommand = &cobra.Command{
	Use:   "isobash",
	Short: "Start an isolated bash process in it's own namespaces.",
	Run: func(cmd *cobra.Command, args []string) {
		jsonNamespaces := `[
			{ "type": "pid" },
			{ "type": "network" },
			{ "type": "ipc" },
			{ "type": "uts" },
			{ "type": "mount" },
			{ "type": "cgroup" }
		]`
		var testNamespaces []specs.LinuxNamespace
		err := json.Unmarshal([]byte(jsonNamespaces), &testNamespaces)
		if err != nil {
			return
		}

		var testNamespaceProfile namespace.ProcNamespaceProfile
		testNamespaceProfile.Namespaces = testNamespaces
		testNamespaceProfile.ProcessBinary = ""

		testNamespaceProfile.StartBashInNewNamespaces()

	},
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	rootCmd.AddCommand(startCommand)

	rootCmd.AddCommand(isoBashCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
