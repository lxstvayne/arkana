package main

import (
	"arkana/lib/variables"
	"arkana/parser"
	"fmt"
)

func main() {
	exampleString := "word = 2 + 2\nword2 = PI + word"
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
	fmt.Printf("%s = %f\n", "word", variables.Get("word"))
	fmt.Printf("%s = %f", "word2", variables.Get("word2"))
}
