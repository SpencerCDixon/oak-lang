# Oak

A simple Ruby/JavaScript like language implemented in Go.

## Why?

This project is just a tool to help me learn about lexers, parsers, and
interpreters/compilers.  **DISCLAIMER**: I wouldn't use this language for
anything serious.  

## What/How?
Oak comes with a [basic CLI](#cli).  To get going quickly:

**Install:**  
```
$ go get github.com/spencercdixon/oak
```

**Execute:**  
```
$ oak run ./path/to/source/code.oak
```

**Explore:**  
```
$ oak repl
```

## Language Features

* Arrays
* Integers
* Strings (double quote)
* Functions (defined with 'fn')
* Closures
* ... more full docs coming ...

## CLI

```
A tree that bears acorns as fruit, and typically has lobed deciduous leaves.  Also, lets you write code.

Usage:
  oak [command]

Available Commands:
  help        Help about any command
  repl        Interactive console to play with oak code.
  run         Execute oak source code
  version     All software needs versions.  This is oaks.

Flags:
      --config string   config file (default is $HOME/.oak.yaml)

Use "oak [command] --help" for more information about a command.
```

## TODO
* [x] add Levenshtein Distance algorithm for better error msgs
* [ ] add ability to write comments in source code
* [ ] add concurrent GET built in function
* [ ] write documentation for basic data structures
* [ ] create some 'real-world' examples of using the language
* [ ] add benchmarking to a real world http example to compare to Ruby
