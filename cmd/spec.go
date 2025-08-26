package cmd

import (
	"encoding/json"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var Spec = &cobra.Command{
	Use:   "spec",
	Short: "Initialize a default config.json file in current working directory.",
	Run: func(cmd *cobra.Command, args []string) {
		spec := utils.GetDefaultConfigJson()
		configJson, err := json.MarshalIndent(spec, "", "  ")
		if err != nil {
			logrus.Fatal("unable to marshal structure 'spec' to string type")
			os.Exit(1)
		}

		err = os.WriteFile("./config.json", []byte(configJson), 0755)
	},
}
