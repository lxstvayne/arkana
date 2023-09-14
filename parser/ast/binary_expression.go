package ast

import "fmt"

type BinaryExpression struct {
	operation rune
	expr1     Expression
	expr2     Expression
}

func NewBinaryExpression(operation rune, expr1 Expression, expr2 Expression) *BinaryExpression {
	return &BinaryExpression{
		operation: operation,
		expr1:     expr1,
		expr2:     expr2,
	}
}

func (expr *BinaryExpression) Eval() float64 {
	switch expr.operation {
	case rune('+'):
		return expr.expr1.Eval() + expr.expr2.Eval()
	case rune('-'):
		return expr.expr1.Eval() - expr.expr2.Eval()
	case rune('*'):
		return expr.expr1.Eval() * expr.expr2.Eval()
	case rune('/'):
		return expr.expr1.Eval() / expr.expr2.Eval()
	default:
		// FIXME
		return expr.expr1.Eval() + expr.expr2.Eval()
	}
}

func (expr *BinaryExpression) String() string {
	return fmt.Sprintf("[%s %c %s]", expr.expr1, expr.operation, expr.expr2)
}
