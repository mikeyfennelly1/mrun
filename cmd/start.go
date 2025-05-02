package cmd

import (
	"encoding/json"
	"github.com/mikeyfennelly1/mrun/src"
	"github.com/mikeyfennelly1/mrun/src/cgroup"
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

		// read contents of ./config.json
		jsonContent, err := os.ReadFile("./config.json")
		if err != nil {
			logrus.Fatalf("error reading config: %v", err)
			return
		}

		// unmarshal byte slice into specs.Spec structure
		err = json.Unmarshal(jsonContent, &spec)
		if err != nil {
			logrus.Fatalf("error creating unmarshalling JSON: %v", err)
			return
		}

		// create a random containerID
		containerID := src.CreateNewContainerID()
		// initialize a new cgroup for the container based on the spec
		err = cgroup.InitCgroup(containerID, spec)
		if err != nil {
			logrus.Fatalf("error initializing cgroup: %v", err)
		}

		err = cgroup.MoveCurrentPidToCgroup(containerID)
		if err != nil {
			logrus.Errorf("error moving process into cgroup: %v", err)
			return
		}

		// set and apply capability sets to the process
		proc.SetAndApplyCapsetToCurrentPid(capability.INHERITABLE, spec.Process.Capabilities.Inheritable)
		proc.SetAndApplyCapsetToCurrentPid(capability.PERMITTED, spec.Process.Capabilities.Permitted)
		proc.SetAndApplyCapsetToCurrentPid(capability.EFFECTIVE, spec.Process.Capabilities.Effective)
		proc.SetAndApplyCapsetToCurrentPid(capability.AMBIENT, spec.Process.Capabilities.Ambient)

		// execs the process with the current process program
		// new program is running in new namespaces in chroot jail.
		err = namespace.RestartInNewNS("chroot")
		if err != nil {
			logrus.Fatalf("unable to enter new namesapces")
			logrus.Exit(1)
		}
	},
}
