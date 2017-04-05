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
* [Error reporting](#errors)
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

## Errors
When running some broken oak code like:  
```
let queue = []

push(queu, 1)
```

Oak will give you nice error messages using [Levenshtien
Distance](https://en.wikipedia.org/wiki/Levenshtein_distance) algorithm.  
```
ERROR: identifier queu not found.

Did you mean one of:

 queue
```

## Development
If you happen to be interested in adding functionality or just want to play with
Oak locally here are some things that might help:

* Go v1.8+ installed
* Check out the [Makefile](./Makefile) for available targets.
  + [I use T.J's mmake for a better Makefile experience](https://github.com/tj/mmake)
  + `go get github.com/tj/mmake/cmd/mmake`
  + Add `alias make=mmake` to bash/zshrc.
  + Now you can run: `make help` to see what targets are available
* Oak uses [Pratt Parsing](http://web.archive.org/web/20151223215421/http://hall.org.ua/halls/wizzard/pdf/Vaughan.Pratt.TDOP.pdf) for its parser.  

## TODO
* [x] add Levenshtein Distance algorithm for better error msgs
* [ ] add ability to write comments in source code
* [ ] add nice HTTP primitives to make oak shine
* [ ] write documentation for basic data structures
* [ ] create some 'real-world' examples of using the language
* [ ] add benchmarking to a real world http example to compare to Ruby
* [ ] add basic module system so code can be separated

