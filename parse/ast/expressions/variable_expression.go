package expressions

import (
	"arkana/lib"
	"arkana/lib/variables"
)

type VariableExpression struct {
	name string
}

func NewVariableExpression(name string) *VariableExpression {
	return &VariableExpression{
		name: name,
	}
}

func (expr *VariableExpression) Eval() lib.Value {
	return variables.Get(expr.name)
}

func (expr *VariableExpression) String() string {
	return expr.name
}
