package main

import (
	"arkana/parser"
	"fmt"
)

func main() {
	exampleString := "(PI + 2) * #f"
	lexer := parser.NewLexer(exampleString)
	tokens := lexer.Tokenize()
	for _, tok := range tokens {
		fmt.Println(tok.TokenType(), string(tok.Text()))
	}
	parser := parser.NewParser(tokens)
	expressions := parser.Parse()
	for _, expr := range expressions {
		fmt.Println(expr, "=", expr.Eval())
	}
}
