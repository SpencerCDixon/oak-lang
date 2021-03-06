package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spencercdixon/oak/evaluator"
	"github.com/spencercdixon/oak/lexer"
	"github.com/spencercdixon/oak/object"
	"github.com/spencercdixon/oak/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// User friendly way to exit the REPL
		if line == "exit" {
			fmt.Println("Exiting REPL... Stay grassy folks.")
			os.Exit(1)
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
