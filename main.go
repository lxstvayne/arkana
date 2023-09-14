package main

import (
	"fmt"
	"main/parser"
)

func main() {
	exampleString := "(#2 + -#2) * #FaBCDEF"
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
