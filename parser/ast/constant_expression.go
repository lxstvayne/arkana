package ast

import (
	"arkana/lib/constants"
)

type ConstantExpression struct {
	name string
}

func NewConstantExpression(name string) *ConstantExpression {
	return &ConstantExpression{
		name: name,
	}
}

func (expr *ConstantExpression) Eval() float64 {
	return constants.Get(expr.name)
}

func (expr *ConstantExpression) String() string {
	return expr.name
}
