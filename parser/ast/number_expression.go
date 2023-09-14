package ast

import "fmt"

type NumberExpression struct {
	number float64
}

func NewNumberExpression(number float64) *NumberExpression {
	return &NumberExpression{
		number: number,
	}
}

func (expr *NumberExpression) Eval() float64 {
	return expr.number
}

func (expr *NumberExpression) String() string {
	return fmt.Sprintf("%f", expr.number)
}
