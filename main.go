package main

import (
	"fmt"
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
var rmCommand = &cobra.Command{
	Use:   "start",
	Short: "Start a container.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	rootCmd.AddCommand(rmCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
