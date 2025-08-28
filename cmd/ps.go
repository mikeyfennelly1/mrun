package cmd

import (
	"fmt"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/cobra"
	"os"
)

var Ps = &cobra.Command{
	Use:   "ps",
	Short: "List all mrun managed containers and corresponding metadata.",
	Run: func(cmd *cobra.Command, args []string) {
		// get a list of the directories in the mrun state dir

		// print the header for the state
		fmt.Printf("%-20s %-20s %-20s %-20s\n", "CONTAINER ID", "VERSION", "STATUS", "BUNDLE LOCATION")

		// find all subdirectory entries
		entries, err := os.ReadDir(state.MrunStateGlobalDirectory)
		if err != nil {
			os.Exit(0)
		}

		for _, e := range entries {
			if e.IsDir() {
				sm := state.GetStateManager(e.Name())
				state, _ := sm.FetchState()
				printFormattedState(state)
			}
		}
	},
}

func printFormattedState(state *specs.State) {
	fmt.Printf("%-20s %-20s %-20s %-20s\n", state.ID, state.Version, state.Status, state.Bundle)
	return
}
