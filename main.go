package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spencercdixon/oak/repl"
)

const NAME = `
_______  _______  ___   _
|       ||   _   ||   | | |
|   _   ||  |_|  ||   |_| |
|  | |  ||       ||      _|
|  |_|  ||       ||     |_
|       ||   _   ||    _  |
|_______||__| |__||___| |_|
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println(NAME)
	fmt.Printf("Hello %s!  This is the Oak programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
