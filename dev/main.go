package main

import (
	"arkana/parse"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./dev/program.txt")

	if err != nil {
		panic(err)
	}

	exampleString := string(data)
	lexer := parse.NewLexer(exampleString)
	tokens := lexer.Tokenize()
	for _, tok := range tokens {
		fmt.Println(tok.TokenType(), string(tok.Text()))
	}
	parser := parse.NewParser(tokens)
	program := parser.Parse()
	fmt.Println(program.String())
	program.Execute()
}
