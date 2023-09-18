package main

import (
	"arkana/parser"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./dev/program.txt")

	if err != nil {
		panic(err)
	}

	exampleString := string(data)
	lexer := parser.NewLexer(exampleString)
	tokens := lexer.Tokenize()
	for _, tok := range tokens {
		fmt.Println(tok.TokenType(), string(tok.Text()))
	}
	parser := parser.NewParser(tokens)
	program := parser.Parse()
	fmt.Println(program.String())
	program.Execute()
}
