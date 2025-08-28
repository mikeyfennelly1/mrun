package cmd

import (
	"fmt"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/spf13/cobra"
)

var State = &cobra.Command{
	Use:   "state",
	Short: "Print the state of a container by container ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerID := args[0]

		if state.StateFileExists(containerID) {
			sm := state.GetStateManager(containerID)
			sm.PrintStateFile()
		} else {
			fmt.Printf("No container with ID: %s\n", containerID)
		}
	},
}
