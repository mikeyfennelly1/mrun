package cmd

import (
	"github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
