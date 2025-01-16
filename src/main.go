package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mrun",
		Short: "A low-level container runtime for managing the lifecycle of OCI compliant Linux containers.",
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
