package cmd

import (
	"github.com/mikeyfennelly1/mrun/src"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		initChain := getInitChain()

		spec, err := src.GetSpec()
		if err != nil {
			logrus.Fatalf(err.Error())
			logrus.Exit(1)
		}

		//TODO decide whether we should get and validate the config outside the init chain or not
		initChain.Execute(spec)

		// execs the process with the current process program
		// new program is running in new namespaces in chroot jail.
		err = src.RestartInNewNS("chroot")
		if err != nil {
			logrus.Fatalf("unable to enter new namesapces")
			logrus.Exit(1)
		}
	},
}

func getInitChain() src.ChainLink {
	// instantiate container state
	parseConfigLink := &src.ParseConfigLink{}

	//TODO init a containerID and pass to InitCgroup. InitCgroup needs to name
	// the cgroup after the containerId.
	// This initContainerState should instantiate an in-memory singleton for the
	// container state to be used by relevant functions, as well as state.json, which
	// is to be updated as init lifecycle progresses.
	initContainerStateLink := &src.InitContainerStateLink{}
	parseConfigLink.SetNext(initContainerStateLink)

	//TODO move the process into control group created from InitCgroupLink.
	initCgroupLink := &src.InitCgroupLink{}
	initContainerStateLink.SetNext(initCgroupLink)

	//TODO pass containerId to function that creates container state directory/file.
	applyCapsetLink := &src.ApplyCapsetLink{}
	initCgroupLink.SetNext(applyCapsetLink)

	return parseConfigLink
}
