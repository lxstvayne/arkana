package main

import (
	"arkana/parser"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./program.txt")

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
	statements := parser.Parse()
	for _, st := range statements {
		fmt.Println(st)
	}
	for _, st := range statements {
		st.Execute()
	}
}
