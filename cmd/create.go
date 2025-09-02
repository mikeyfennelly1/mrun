package cmd

import (
	"os"

	"github.com/mikeyfennelly1/mrun/init/libinitsteps"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Create is a standard command in an OCI container runtime CLI interface
// that kicks off step 1 of the OCI container lifecycle.
// Create initializes the environment that we call a "container", but does
// not start the application itself.
// It does however, delegate the starting of the container's application
// to the runtime start command, passing the container ID
// as an argument to that command.
//
// See the OCI lifecycle overview: https://github.com/opencontainers/runtime-spec/blob/main/runtime.md#lifecycle
// See the create command: https://github.com/opencontainers/runtime-spec/blob/main/runtime.md#create
var Create = &cobra.Command{
	Use:   "create",
	Short: "Create a container based on an OCI bundle.",
	Args:  cobra.ExactArgs(1), // path to the bundle
	Run: func(cmd *cobra.Command, args []string) {
		bundlePath := args[0]

		// container is not being created
		// get config.json in pwd
		spec, err := utils.GetConfigJson(bundlePath)
		if err != nil {
			panic(err)
		}
		stepNames := [...]string{
			"init-cgroup",
			"capset",
		}
		// initialize container sm
		containerID := utils.NewContainerID()
		sm, err := state.NewContainerState(containerID)
		if err != nil {
			logrus.Fatal("unable to initialize container state manager")
		}

		for _, stepName := range stepNames {
			step, err := libinitsteps.StepFactory(stepName)
			if err != nil {
				sm.CleanUp()
				logrus.Fatalf("unable to get init step %s from StepFactory: %v", stepName, err)
				os.Exit(1)
			}
			err = step.Execute(spec, sm)
			if err != nil {
				logrus.Errorf("mrun create failed with error: %v", err)
				sm.CleanUp()
				os.Exit(1)
			}
		}

		startInNewNSStep, err := libinitsteps.StepFactory("namespace")
		startInNewNSStep.Execute(spec, sm)
		if err != nil {
			sm.CleanUp()
			logrus.Fatalf("unable to get init step %s from StepFactory: %v", "namespace", err)
		}
	},
}
