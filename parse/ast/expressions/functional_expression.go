package expressions

import (
	"arkana/lib"
	"arkana/lib/functions"
	"arkana/lib/variables"
	"fmt"
)

type FunctionalExpression struct {
	name string
	args []lib.Expression
}

func NewFunctionalExpression(name string, args []lib.Expression) *FunctionalExpression {
	return &FunctionalExpression{
		name: name,
		args: args,
	}
}

func (expr *FunctionalExpression) AddArgument(arg lib.Expression) {
	expr.args = append(expr.args, arg)
}

func (expr *FunctionalExpression) Eval() lib.Value {
	values := []lib.Value{}
	for _, arg := range expr.args {
		values = append(values, arg.Eval())
	}

	f := functions.Get(expr.name)

	switch f.(type) {
	case *lib.DefineFunction:
		defineFunc := f.(*lib.DefineFunction)
		if len(expr.args) != defineFunc.ArgsCount() {
			panic("Args count mismatch")
		}

		// FIXME
		for i := 0; i < len(expr.args); i++ {
			variables.Set(defineFunc.ArgName(i), values[i])
		}
	}

	return f.Execute(values)
}

func (expr *FunctionalExpression) String() string {
	return fmt.Sprintf("Function name:%s", expr.name)
}
