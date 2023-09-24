package ast

import (
	"arkana/lib"
	"arkana/lib/functions"
	"fmt"
)

type FunctionalExpression struct {
	name string
	args []Expression
}

func NewFunctionalExpression(name string, args []Expression) *FunctionalExpression {
	return &FunctionalExpression{
		name: name,
		args: args,
	}
}

func (expr *FunctionalExpression) AddArgument(arg Expression) {
	expr.args = append(expr.args, arg)
}

func (expr *FunctionalExpression) Eval() lib.Value {
	values := []lib.Value{}
	for _, arg := range expr.args {
		values = append(values, arg.Eval())
	}

	return functions.Get(expr.name).Execute(values...)
}

func (expr *FunctionalExpression) String() string {
	return fmt.Sprintf("Function name:%s", expr.name)
}
