package cmd

import (
	"github.com/mikeyfennelly1/mrun/src"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/syndtr/gocapability/capability"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		var spec specs.Spec

		config, err := src.parseConfig("./config.json")
		if err != nil {
			logrus.Fatalf("could not find file ./config.json")
			logrus.Exit(1)
		}
		// create a random containerID
		containerID := src.NewContainerID()
		// initialize a new cgroup for the container based on the spec
		err = src.InitCgroup(containerID, spec)
		if err != nil {
			logrus.Fatalf("error initializing cgroup: %v", err)
		}

		err = src.MoveCurrentPidToCgroup(containerID)
		if err != nil {
			logrus.Errorf("error moving process into cgroup: %v", err)
			return
		}

		// set and apply capability sets to the process
		src.SetAndApplyCapsetToCurrentPid(capability.INHERITABLE, spec.Process.Capabilities.Inheritable)
		src.SetAndApplyCapsetToCurrentPid(capability.PERMITTED, spec.Process.Capabilities.Permitted)
		src.SetAndApplyCapsetToCurrentPid(capability.EFFECTIVE, spec.Process.Capabilities.Effective)
		src.SetAndApplyCapsetToCurrentPid(capability.AMBIENT, spec.Process.Capabilities.Ambient)

		err = src.InitContainerStateDirAndFile(containerID, spec)
		if err != nil {
			logrus.Errorf("could not intialize container state: %v", err)
			return
		}

		// execs the process with the current process program
		// new program is running in new namespaces in chroot jail.
		err = src.RestartInNewNS("chroot")
		if err != nil {
			logrus.Fatalf("unable to enter new namesapces")
			logrus.Exit(1)
		}
	},
}

func getInitChain() src.ChainItem {
	// instantiate container state
	parseConfigLink := &src.ParseConfigLink{}

	parseConfigLink.SetNext()

	return parseConfigLink
}
