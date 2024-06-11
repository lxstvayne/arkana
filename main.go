package main

import (
	"arkana/parse"
	"flag"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 1 {
		readAndRunProgram(args[0])
	} else {
		panic("Unknown usage")
	}
}

func readAndRunProgram(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	programString := string(data)
	lexer := parse.NewLexer(programString)
	tokens := lexer.Tokenize()
	parser := parse.NewParser(tokens)
	program := parser.Parse()
	program.Execute()
}
