package cmd

import (
	"fmt"

	"github.com/mikeyfennelly1/mrun/init/libinitsteps"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/spf13/cobra"
)

// Start triggers step 6 of the OCI container lifecycle.
//
// A misconception is that the start command has no effect on the environment
// that the application process resides in. This step in the process does have
// an effect on what the container actually sees and does, just before it starts:
//
// 1. Changes the filesystem root of the environment (with the chroot syscall)
// 2. Sets the users, and groups
// 3. Sets the environment variables that the containerized application will run with.
// 4. Sets the hostname of the container.
// 5. Creates filesystem.
// 6. Executes the application binary.
//
// See https://github.com/opencontainers/runtime-spec/blob/main/runtime.md#start
var Start = &cobra.Command{
	Use:   "start",
	Short: "Start a containerized application from an already created container.",
	Args:  cobra.ExactArgs(1), // Container ID
	Run: func(cmd *cobra.Command, args []string) {
		containerID := args[0]

		sm := state.GetStateManager(containerID)
		state, err := sm.FetchState()
		if err != nil {
			panic(err)
		}

		spec, err := utils.GetConfigJson(fmt.Sprintf("%s/config.json", state.Bundle))
		if err != nil {
			panic(err)
		}

		stepNames := [...]string{
			"chroot",
			"rlimit",
			"set-env",
			"hostname",
			"create-fs",
			"exec-bin",
		}
		for _, stepName := range stepNames {
			step, err := libinitsteps.StepFactory(stepName)
			if err != nil {
				panic(err)
			}
			err = step.Execute(spec, sm)
			if err != nil {
				panic(err)
			}
		}
	},
}
