package cmd

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/cobra"
)

var Ps = &cobra.Command{
	Use:   "ps",
	Short: "List all mrun managed containers and corresponding metadata.",
	Run: func(cmd *cobra.Command, args []string) {
		// get a list of the directories in the mrun state dir

		// print the header for the state
		// loop through each directory
		// unmarshal <directory>/state.json to specs.State struct
		// pass state structure to printing function
		fmt.Printf("%-20s %-20s %-20s %-20s\n", "Container ID", "Version", "Status", "Bundle Location")
		testSpec1 := specs.State{
			ID:      "XYXYXY",
			Version: "latest",
			Status:  "created",
			Bundle:  "/non/existent/path",
		}
		printFormattedState(&testSpec1)
	},
}

func printFormattedState(state *specs.State) {
	fmt.Printf("%-20s %-20s %-20s %-20s\n", state.ID, state.Version, state.Status, state.Bundle)
	return
}
