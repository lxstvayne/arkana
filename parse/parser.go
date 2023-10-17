package parse

import (
	"arkana/lib"
	"arkana/parse/ast/expressions"
	"arkana/parse/ast/statements"
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
func (parser *Parser) Parse() lib.Statement {
	result := statements.NewBlockStatement(nil)

	for {
		if parser.match(TOKENTYPE_EOF) {
			break
		}

		result.Add(parser.statement())
	}

	return result
}

func (parser *Parser) block() lib.Statement {
	block := statements.NewBlockStatement(nil)
	parser.consume(TOKENTYPE_LBRACE)
	for {
		if parser.match(TOKENTYPE_RBRACE) {
			break
		}

		block.Add(parser.statement())
	}
	return block
}

func (parser *Parser) statementOrBlock() lib.Statement {
	if parser.get(0).TokenType() == TOKENTYPE_LBRACE {
		return parser.block()
	} else {
		return parser.statement()
	}
}

func (parser *Parser) statement() lib.Statement {
	if parser.match(TOKENTYPE_IF) {
		return parser.ifElseStatement()
	}
	if parser.match(TOKENTYPE_WHILE) {
		return parser.whileStatement()
	}
	if parser.match(TOKENTYPE_FOR) {
		return parser.forStatement()
	}
	if parser.match(TOKENTYPE_BREAK) {
		return statements.NewBreakStatement()
	}
	if parser.match(TOKENTYPE_CONTINUE) {
		return statements.NewContinueStatement()
	}
	if parser.match(TOKENTYPE_FUNC) {
		return parser.functionDefine()
	}
	if parser.get(0).TokenType() == TOKENTYPE_WORD && parser.get(1).TokenType() == TOKENTYPE_LPAR {
		return statements.NewFunctionStatement(parser.function())
	}
	return parser.assignmentStatement()
}

func (parser *Parser) assignmentStatement() lib.Statement {
	// WORD EQ
	current := parser.get(0)
	if parser.match(TOKENTYPE_WORD) && parser.get(0).TokenType() == TOKENTYPE_EQ {
		variable := current.Text()
		parser.consume(TOKENTYPE_EQ)
		return statements.NewAssignmentStatement(variable, parser.expression())
	}

	panic("Unknown statement")
}

func (parser *Parser) ifElseStatement() lib.Statement {
	condition := parser.expression()
	ifStmt := parser.statementOrBlock()
	var elseStmt lib.Statement

	if parser.match(TOKENTYPE_ELSE) {
		elseStmt = parser.statementOrBlock()
	}

	return statements.NewIfStatement(condition, ifStmt, elseStmt)
}

func (parser *Parser) whileStatement() lib.Statement {
	condition := parser.expression()
	stmt := parser.statementOrBlock()
	return statements.NewWhileStatement(condition, stmt)
}

func (parser *Parser) forStatement() lib.Statement {
	initialization := parser.assignmentStatement()
	parser.consume(TOKENTYPE_COMMA)
	termination := parser.expression()
	parser.consume(TOKENTYPE_COMMA)
	increment := parser.assignmentStatement()
	stmt := parser.statementOrBlock()
	return statements.NewForStatement(initialization, termination, increment, stmt)
}

func (parser *Parser) functionDefine() *statements.FunctionDefineStatement {
	name := string(parser.consume(TOKENTYPE_WORD).Text())
	parser.consume(TOKENTYPE_LPAR)
	argNames := []string{}
	for {
		if parser.match(TOKENTYPE_RPAR) {
			break
		}

		argNames = append(argNames, string(parser.consume(TOKENTYPE_WORD).Text()))
		parser.match(TOKENTYPE_COMMA)
	}

	body := parser.statementOrBlock()

	return statements.NewFunctionDefineStatement(name, argNames, body)
}

func (parser *Parser) function() *expressions.FunctionalExpression {
	name := string(parser.consume(TOKENTYPE_WORD).Text())
	parser.consume(TOKENTYPE_LPAR)
	function := expressions.NewFunctionalExpression(name, nil)
	for {
		if parser.match(TOKENTYPE_RPAR) {
			break
		}

		function.AddArgument(parser.expression())
		parser.match(TOKENTYPE_COMMA)
	}

	return function
}

func (parser *Parser) expression() lib.Expression {
	return parser.logicalOr()
}

func (parser *Parser) logicalOr() lib.Expression {
	result := parser.logicalAnd()
	for {
		if parser.match(TOKENTYPE_BARBAR) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_OR, result, parser.logicalAnd())
			continue
		}
		break
	}

	return result
}
func (parser *Parser) logicalAnd() lib.Expression {
	result := parser.equality()

	for {
		if parser.match(TOKENTYPE_AMPAMP) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_AND, result, parser.equality())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) equality() lib.Expression {
	result := parser.conditional()

	if parser.match(TOKENTYPE_EQEQ) {
		return expressions.NewConditionalExpression(expressions.OPERATOR_EQUALS, result, parser.conditional())
	}
	if parser.match(TOKENTYPE_EXCLEQ) {
		return expressions.NewConditionalExpression(expressions.OPERATOR_NOT_EQUALS, result, parser.conditional())
	}

	return result
}

func (parser *Parser) conditional() lib.Expression {
	result := parser.additive()

	for {
		if parser.match(TOKENTYPE_LT) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_LT, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_LTEQ) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_LTEQ, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_GT) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_GT, result, parser.additive())
			continue
		}
		if parser.match(TOKENTYPE_GTEQ) {
			result = expressions.NewConditionalExpression(expressions.OPERATOR_GTEQ, result, parser.additive())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) additive() lib.Expression {
	result := parser.multiplicative()

	for {
		if parser.match(TOKENTYPE_PLUS) {
			result = expressions.NewBinaryExpression(rune('+'), result, parser.multiplicative())
			continue
		}
		if parser.match(TOKENTYPE_MINUS) {
			result = expressions.NewBinaryExpression(rune('-'), result, parser.multiplicative())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) multiplicative() lib.Expression {
	result := parser.unary()

	for {
		if parser.match(TOKENTYPE_STAR) {
			result = expressions.NewBinaryExpression(rune('*'), result, parser.unary())
			continue
		}
		if parser.match(TOKENTYPE_SLASH) {
			result = expressions.NewBinaryExpression(rune('/'), result, parser.unary())
			continue
		}
		break
	}

	return result
}

func (parser *Parser) unary() lib.Expression {
	if parser.match(TOKENTYPE_MINUS) {
		return expressions.NewUnaryExpression(rune('-'), parser.primary())
	} else if parser.match(TOKENTYPE_PLUS) {
		return parser.primary()
	}
	return parser.primary()
}

func (parser *Parser) primary() lib.Expression {
	currentTok := parser.get(0)
	if parser.match(TOKENTYPE_NUMBER) {
		// No handle error?
		number, err := strconv.ParseFloat(string(currentTok.Text()), 32)
		if err != nil {
			panic(err)
		}
		expr := expressions.NewValueExpression(number)
		return expr
	}
	if parser.get(0).TokenType() == TOKENTYPE_WORD && parser.get(1).TokenType() == TOKENTYPE_LPAR {
		return parser.function()
	}
	if parser.match(TOKENTYPE_WORD) {
		return expressions.NewVariableExpression(string(currentTok.Text()))
	}
	if parser.match(TOKENTYPE_TEXT) {
		return expressions.NewValueExpression(string(currentTok.Text()))
	}
	if parser.match(TOKENTYPE_HEX_NUMBER) {
		// No handle error?
		number, err := strconv.ParseUint(string(currentTok.Text()), 16, 32)
		if err != nil {
			panic(err)
		}
		expr := expressions.NewValueExpression(float64(number))
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
