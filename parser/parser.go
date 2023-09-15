package parser

import (
	"arkana/parser/ast"
	"fmt"
	"strconv"
)

type Parser struct {
	tokens   []*Token
	position int
	size     int
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{
		tokens: tokens,
		size:   len(tokens),
	}
}

// Метод рекурсивного спуска
func (parser *Parser) Parse() (result []ast.Statement) {
	for {
		if parser.match(TOKENTYPE_EOF) {
			return
		}

		result = append(result, parser.statement())
	}
}

func (parser *Parser) statement() ast.Statement {
	return parser.assignmentStatement()
}

func (parser *Parser) assignmentStatement() ast.Statement {
	// WORD EQ
	current := parser.get(0)
	if parser.match(TOKENTYPE_WORD) && parser.get(0).TokenType() == TOKENTYPE_EQ {
		variable := current.Text()
		parser.consume(TOKENTYPE_EQ)
		return ast.NewAssignmentStatement(variable, parser.expression())
	}

	panic("Unknown statement")
}

func (parser *Parser) expression() ast.Expression {
	return parser.additive()
}

func (parser *Parser) additive() ast.Expression {
	result := parser.multiplicative()

	for {
		if parser.match(TOKENTYPE_PLUS) {
			result = ast.NewBinaryExpression(rune('+'), result, parser.multiplicative())
			continue
		}
		if parser.match(TOKENTYPE_MINUS) {
			result = ast.NewBinaryExpression(rune('-'), result, parser.multiplicative())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) multiplicative() ast.Expression {
	result := parser.unary()

	for {
		if parser.match(TOKENTYPE_STAR) {
			result = ast.NewBinaryExpression(rune('*'), result, parser.unary())
			continue
		}
		if parser.match(TOKENTYPE_SLASH) {
			result = ast.NewBinaryExpression(rune('/'), result, parser.unary())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) unary() ast.Expression {
	if parser.match(TOKENTYPE_MINUS) {
		return ast.NewUnaryExpression(rune('-'), parser.primary())
	} else if parser.match(TOKENTYPE_PLUS) {
		return parser.primary()
	}
	return parser.primary()
}

func (parser *Parser) primary() ast.Expression {
	currentTok := parser.get(0)
	if parser.match(TOKENTYPE_NUMBER) {
		// No handle error?
		number, err := strconv.ParseFloat(string(currentTok.Text()), 32)
		if err != nil {
			panic(err)
		}
		expr := ast.NewNumberExpression(number)
		return expr
	}
	if parser.match(TOKENTYPE_WORD) {
		return ast.NewConstantExpression(string(currentTok.Text()))
	}
	if parser.match(TOKENTYPE_HEX_NUMBER) {
		// No handle error?
		number, err := strconv.ParseUint(string(currentTok.Text()), 16, 32)
		if err != nil {
			panic(err)
		}
		expr := ast.NewNumberExpression(float64(number))
		return expr
	}
	if parser.match(TOKENTYPE_LPAR) {
		result := parser.expression()
		parser.match(TOKENTYPE_RPAR)
		return result
	}
	panic("Unknown Expression")
}

func (parser *Parser) consume(tokenType TokenType) *Token {
	currentTok := parser.get(0)
	if tokenType != currentTok.TokenType() {
		panic(fmt.Sprintf("Token %s doesnt match %s", string(currentTok.Text()), tokenType))
	}
	parser.position += 1
	return currentTok
}

func (parser *Parser) match(tokenType TokenType) bool {
	currentTok := parser.get(0)
	if tokenType != currentTok.TokenType() {
		return false
	}
	parser.position += 1
	return true
}

func (parser *Parser) get(relativePos int) *Token {
	pos := parser.position + relativePos
	if pos >= parser.size {
		return TOKEN_EOF
	}
	return parser.tokens[pos]
}
