package ast

import "fmt"

type UnaryExpression struct {
	operation rune
	expr      Expression
}

func NewUnaryExpression(operation rune, expr Expression) *UnaryExpression {
	return &UnaryExpression{
		operation: operation,
		expr:      expr,
	}
}

func (expr *UnaryExpression) Eval() float64 {
	switch expr.operation {
	case rune('-'):
		return -expr.expr.Eval()
	case rune('+'):
		return expr.expr.Eval()
	default:
		panic("Unary expr!")
	}
}

func (expr *UnaryExpression) String() string {
	return fmt.Sprintf("%c %s", expr.operation, expr.expr)
}
