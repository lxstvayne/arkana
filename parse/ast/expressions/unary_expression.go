package expressions

import (
	"arkana/lib"
	"fmt"
)

type UnaryExpression struct {
	operation rune
	expr      lib.Expression
}

func NewUnaryExpression(operation rune, expr lib.Expression) *UnaryExpression {
	return &UnaryExpression{
		operation: operation,
		expr:      expr,
	}
}

func (expr *UnaryExpression) Eval() lib.Value {
	switch expr.operation {
	case rune('-'):
		return lib.NewNumberValue(-expr.expr.Eval().Float64())
	case rune('+'):
		return expr.expr.Eval()
	default:
		panic("Unary expr!")
	}
}

func (expr *UnaryExpression) String() string {
	return fmt.Sprintf("%c %s", expr.operation, expr.expr)
}
