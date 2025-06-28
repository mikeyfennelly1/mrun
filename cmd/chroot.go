package cmd

import (
	"github.com/mikeyfennelly1/mrun/src"
	"github.com/spf13/cobra"
)

var Chroot = &cobra.Command{
	Use:   "chroot",
	Short: "change the root for the binary.",
	Run: func(cmd *cobra.Command, args []string) {
		postChrootInitChain := getPostChrootInitChain()

		spec, err := src.GetSpec()
		if err != nil {
			return
		}

		postChrootInitChain.Execute(spec)
	},
}

func getPostChrootInitChain() src.ChainLink {
	chrootLink := src.ChrootLink{}

	changeProcRootLink := src.ChangeProcessDirToNewRootLink{}
	chrootLink.SetNext(changeProcRootLink)

	setUsersAndGroupsLink := src.SetUsersAndGroupsLink{}
	changeProcRootLink.SetNext(setUsersAndGroupsLink)

	setRlimitLink := src.SetRLIMITLink{}
	setUsersAndGroupsLink.SetNext(setUsersAndGroupsLink)

	setEnvVarsLink := src.SetEnvVarsLink{}
	setRlimitLink.SetNext(setEnvVarsLink)

	createFileSystemLink := src.CreateFileSystemLink{}
	setEnvVarsLink.SetNext(createFileSystemLink)

	execBinaryLink := src.ExecBinaryLink{}
	createFileSystemLink.SetNext(execBinaryLink)

	return &chrootLink
}
