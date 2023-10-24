package expressions

import (
	"arkana/lib"
	"arkana/lib/variables"
	"fmt"
)

type ArrayAccessExpression struct {
	variable string
	idx      lib.Expression
}

func NewArrayAccessExpression(variable string, idx lib.Expression) *ArrayAccessExpression {
	return &ArrayAccessExpression{
		variable: variable,
		idx:      idx,
	}
}

func (expr *ArrayAccessExpression) Eval() lib.Value {
	variable := variables.Get(expr.variable)
	switch arr := variable.(type) {
	case *lib.ArrayValue:
		return arr.Get(int(expr.idx.Eval().Float64()))
	default:
		panic("Not array value")
	}
}

func (expr *ArrayAccessExpression) String() string {
	return fmt.Sprintf("%s[%s]", expr.variable, expr.idx)
}
