package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "All software needs versions.  This is oaks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Oak version: %s\n", VERSION)
	},
}
