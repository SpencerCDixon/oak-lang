package cmd

import (
	"github.com/spencercdixon/oak/server"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts an HTTP Server for debugging oak code.",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}
