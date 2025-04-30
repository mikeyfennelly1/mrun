package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mikeyfennelly1/mrun/src/namespace"
	"github.com/mikeyfennelly1/mrun/src/proc"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/syndtr/gocapability/capability"
	"os"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		var spec specs.Spec
		jsonContent, err := os.ReadFile("./config.json")
		if err != nil {
			fmt.Printf("error reading config: %v", err)
			return
		}
		err = json.Unmarshal(jsonContent, &spec)
		if err != nil {
			fmt.Printf("error creating unmarshalling JSON: %v", err)
			return
		}

		proc.SetAndApplyCapsetToCurrentPid(capability.INHERITABLE, spec.Process.Capabilities.Inheritable)

		proc.SetAndApplyCapsetToCurrentPid(capability.PERMITTED, spec.Process.Capabilities.Permitted)

		proc.SetAndApplyCapsetToCurrentPid(capability.EFFECTIVE, spec.Process.Capabilities.Effective)

		proc.SetAndApplyCapsetToCurrentPid(capability.AMBIENT, spec.Process.Capabilities.Ambient)

		err = namespace.RestartInNewNS("chroot")
		if err != nil {
			logrus.Fatalf("unable to enter new namesapces")
			logrus.Exit(1)
		}
	},
}
