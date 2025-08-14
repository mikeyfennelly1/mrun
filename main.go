package main

import (
	"fmt"
	"github.com/mikeyfennelly1/mrun/cmd"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mrun", // The name of the command
	Short: "A low-level container runtime.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // ignoring this error
	},
}

func main() {
	rootCmd.AddCommand(cmd.Start)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
