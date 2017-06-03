package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spencercdixon/oak/ast"
	"github.com/spencercdixon/oak/evaluator"
	"github.com/spencercdixon/oak/lexer"
	"github.com/spencercdixon/oak/object"
	"github.com/spencercdixon/oak/parser"
)

type OakRequest struct {
	Program string `json="program"`
}

// renderJson is a utility function for responding to HTTP requests with
// marshalled JSON.
func renderJson(w http.ResponseWriter, status int, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

// generateAst takes a string oak program and returns the AST generated for that
// program.
func generateAst(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func astHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Missing request body", http.StatusBadRequest)
			return
		}

		req := OakRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		program := generateAst(req.Program)
		renderJson(w, 200, program)
	} else {
		http.Error(w, "Unknown method", http.StatusBadRequest)
		return
	}
}

func evaluationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Missing request body", http.StatusBadRequest)
			return
		}

		req := OakRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		env := object.NewEnvironment()
		program := generateAst(req.Program)
		evaluated := evaluator.Eval(program, env)
		renderJson(w, 200, evaluated)
	} else {
		http.Error(w, "Unknown method", http.StatusBadRequest)
		return
	}
}

func Start() {
	// Set up handlers
	http.HandleFunc("/ast", astHandler)
	http.HandleFunc("/evaluate", evaluationHandler)

	// Listen to server
	fmt.Println("Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
