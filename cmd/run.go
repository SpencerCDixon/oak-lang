package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spencercdixon/oak/evaluator"
	"github.com/spencercdixon/oak/lexer"
	"github.com/spencercdixon/oak/object"
	"github.com/spencercdixon/oak/parser"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Execute oak source code",
	Run: func(cmd *cobra.Command, args []string) {
		if !fileExists(args[0]) {
			fmt.Printf("File %s doesn't exist\n", args[0])
			os.Exit(-1)
		}

		// Read source
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
		source := string(file)

		// Create a new environment for execution
		environment := object.NewEnvironment()

		// Lex, Parse, & Evaluate
		l := lexer.New(source)
		p := parser.New(l)
		program := p.ParseProgram()
		result := evaluator.Eval(program, environment)

		if result != nil {
			os.Stdin.WriteString(result.Inspect())
			os.Stdin.WriteString("\n")
		}
	},
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
