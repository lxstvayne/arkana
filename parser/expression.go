package parser

import "fmt"

type Expression interface {
	Eval() float32
}

type NumberExpression float32

func (ne *NumberExpression) Eval() float32 {
	return float32(*ne)
}

func (ne *NumberExpression) String() string {
	number := float32(*ne)
	return fmt.Sprintf("%f", number)
}

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

func (be *BinaryExpression) Eval() float32 {
	switch be.operation {
	case rune('+'):
		return be.expr1.Eval() + be.expr2.Eval()
	case rune('-'):
		return be.expr1.Eval() - be.expr2.Eval()
	case rune('*'):
		return be.expr1.Eval() * be.expr2.Eval()
	case rune('/'):
		return be.expr1.Eval() / be.expr2.Eval()
	default:
		// FIXME
		return be.expr1.Eval() + be.expr2.Eval()
	}
}

func (be *BinaryExpression) String() string {
	return fmt.Sprintf("%s %c %s", be.expr1, be.operation, be.expr2)
}

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

func (ue *UnaryExpression) Eval() float32 {
	switch ue.operation {
	case rune('-'):
		return -ue.expr.Eval()
	case rune('+'):
		return ue.expr.Eval()
	default:
		panic("Unary expr!")
	}
}

func (ue *UnaryExpression) String() string {
	return fmt.Sprintf("%c %s", ue.operation, ue.expr)
}
