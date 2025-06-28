package cmd

import (
	"github.com/mikeyfennelly1/mrun/container"
	"github.com/spf13/cobra"
)

var Chroot = &cobra.Command{
	Use:   "chroot",
	Short: "change the root for the binary.",
	Run: func(cmd *cobra.Command, args []string) {
		postChrootInitChain := getPostChrootInitChain()

		spec, err := container.GetSpec()
		if err != nil {
			return
		}

		postChrootInitChain.Execute(spec)
	},
}

func getPostChrootInitChain() container.ChainLink {
	chrootLink := container.ChrootLink{}

	changeProcRootLink := container.ChangeProcessDirToNewRootLink{}
	chrootLink.SetNext(changeProcRootLink)

	setUsersAndGroupsLink := container.SetUsersAndGroupsLink{}
	changeProcRootLink.SetNext(setUsersAndGroupsLink)

	setRlimitLink := container.SetRLIMITLink{}
	setUsersAndGroupsLink.SetNext(setUsersAndGroupsLink)

	setEnvVarsLink := container.SetEnvVarsLink{}
	setRlimitLink.SetNext(setEnvVarsLink)

	createFileSystemLink := container.CreateFileSystemLink{}
	setEnvVarsLink.SetNext(createFileSystemLink)

	execBinaryLink := container.ExecBinaryLink{}
	createFileSystemLink.SetNext(execBinaryLink)

	return &chrootLink
}
