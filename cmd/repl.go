package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spencercdixon/oak/repl"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(replCommand)
}

const NAME = `
_______  _______  ___   _
|       ||   _   ||   | | |
|   _   ||  |_|  ||   |_| |
|  | |  ||       ||      _|
|  |_|  ||       ||     |_
|       ||   _   ||    _  |
|_______||__| |__||___| |_|
`

var replCommand = &cobra.Command{
	Use:   "repl",
	Short: "Interactive console to play with oak code.",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Println(NAME)
		fmt.Printf("Hello %s!  This is the Oak programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands ('exit' to leave)\n")
		repl.Start(os.Stdin, os.Stdout)
	},
}
