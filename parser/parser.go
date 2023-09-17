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
	if parser.match(TOKENTYPE_PRINT) {
		return ast.NewPrintStatement(parser.expression())
	}
	if parser.match(TOKENTYPE_IF) {
		return parser.ifElseStatement()
	}
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

func (parser *Parser) ifElseStatement() ast.Statement {
	condition := parser.expression()
	ifStmt := parser.statement()
	var elseStmt ast.Statement

	if parser.match(TOKENTYPE_ELSE) {
		elseStmt = parser.statement() // ?
	}

	return ast.NewIfStatement(condition, ifStmt, elseStmt)
}

func (parser *Parser) expression() ast.Expression {
	return parser.logicalOr()
}

func (parser *Parser) logicalOr() ast.Expression {
	result := parser.logicalAnd()
	for {
		if parser.match(TOKENTYPE_BARBAR) {
			result = ast.NewConditionalExpression(ast.OPERATOR_OR, result, parser.logicalAnd())
			continue
		}
		break
	}

	return result
}
func (parser *Parser) logicalAnd() ast.Expression {
	result := parser.equality()

	for {
		if parser.match(TOKENTYPE_AMPAMP) {
			result = ast.NewConditionalExpression(ast.OPERATOR_AND, result, parser.equality())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) equality() ast.Expression {
	result := parser.conditional()

	if parser.match(TOKENTYPE_EQEQ) {
		return ast.NewConditionalExpression(ast.OPERATOR_EQUALS, result, parser.conditional())
	}
	if parser.match(TOKENTYPE_EXCLEQ) {
		return ast.NewConditionalExpression(ast.OPERATOR_NOT_EQUALS, result, parser.conditional())
	}

	return result
}

func (parser *Parser) conditional() ast.Expression {
	result := parser.additive()

	for {
		if parser.match(TOKENTYPE_LT) {
			result = ast.NewConditionalExpression(ast.OPERATOR_LT, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_LTEQ) {
			result = ast.NewConditionalExpression(ast.OPERATOR_LTEQ, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_GT) {
			result = ast.NewConditionalExpression(ast.OPERATOR_GT, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_GTEQ) {
			result = ast.NewConditionalExpression(ast.OPERATOR_GTEQ, result, parser.additive())
			continue
		}
		break
	}

	return result
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
		expr := ast.NewValueExpression(number)
		return expr
	}
	if parser.match(TOKENTYPE_WORD) {
		return ast.NewVariableExpression(string(currentTok.Text()))
	}
	if parser.match(TOKENTYPE_TEXT) {
		return ast.NewValueExpression(string(currentTok.Text()))
	}
	if parser.match(TOKENTYPE_HEX_NUMBER) {
		// No handle error?
		number, err := strconv.ParseUint(string(currentTok.Text()), 16, 32)
		if err != nil {
			panic(err)
		}
		expr := ast.NewValueExpression(float64(number))
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
