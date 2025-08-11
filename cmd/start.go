package cmd

import (
	"github.com/mikeyfennelly1/mrun/init"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		hub, err := init.Init()
		if err != nil {
			logrus.Fatal("could not initialize hub")
		}
		cur := hub.GetSteps()
		for {
			if cur != nil {
				cur.Execute()

				cur = *cur.Next()
			}
			break
		}
	},
}
