package cmd

import (
	"github.com/mikeyfennelly1/mrun/init/libinitsteps"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/spf13/cobra"
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

			for _, stepName := range stepNames {
				step, err := libinitsteps.StepFactory(stepName)
				if err != nil {
					panic(err)
				}
				err = step.Execute(spec)
				if err != nil {
					panic(err)
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
				err = step.Execute(spec)
				if err != nil {
					panic(err)
				}
			}
		}
	},
}
