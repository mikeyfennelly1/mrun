package cmd

import (
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/spf13/cobra"
	"os"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		var containerIsAlreadyBeingCreated bool = false
		if !containerIsAlreadyBeingCreated {
			// pre chroot steps
			_, err := utils.GetConfigJson("./config.json")
			if err != nil {
				os.Exit(1)
			}

		} else {
			// post chroot steps
		}
	},
}
