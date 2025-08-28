package cmd

import (
	"github.com/mikeyfennelly1/mrun/init/libinitsteps"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start an isolated environment.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// container is not being created
			// get config.json in pwd
			spec, err := utils.GetConfigJson("./config.json")
			if err != nil {
				panic(err)
			}
			stepNames := [...]string{
				"init-cgroup",
				"capset",
				"namespace",
			}
			// initialize container sm
			containerID := utils.NewContainerID()
			sm, err := state.NewContainerState(containerID)
			if err != nil {
				logrus.Fatal("unable to initialize container sm")
				panic("exiting...")
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
					logrus.Errorf("mrun start failed with error: %v", err)
					sm.CleanUp()
					os.Exit(1)
				}
			}

		} else {
			// get config.json in pwd
			spec, err := utils.GetConfigJson("./config.json")
			if err != nil {
				panic(err)
			}
			stepNames := [...]string{
				"chroot",
				"usersgroups",
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
				err = step.Execute(spec, nil)
				if err != nil {
					panic(err)
				}
			}
		}
	},
}
