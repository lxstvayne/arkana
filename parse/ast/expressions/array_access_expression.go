package expressions

import (
	"arkana/lib"
	"arkana/lib/variables"
	"fmt"
)

type ArrayAccessExpression struct {
	variable string
	indexes  []lib.Expression
}

func NewArrayAccessExpression(variable string, indexes []lib.Expression) *ArrayAccessExpression {
	return &ArrayAccessExpression{
		variable: variable,
		indexes:  indexes,
	}
}

func (expr *ArrayAccessExpression) Eval() lib.Value {
	return expr.getArray().Get(expr.lastIndex())
}

func (expr *ArrayAccessExpression) String() string {
	return fmt.Sprintf("%s[%s]", expr.variable, string(expr.lastIndex()))
}

func (expr *ArrayAccessExpression) getArray() *lib.ArrayValue {
	array := expr.consumeArray(variables.Get(expr.variable))
	last := len(expr.indexes) - 1
	for i := 0; i < last; i++ {
		array = expr.consumeArray(array.Get(expr.index(i)))
	}

	return array
}

func (expr *ArrayAccessExpression) lastIndex() int {
	return expr.index(len(expr.indexes) - 1)
}

func (expr *ArrayAccessExpression) index(i int) int {
	return int(expr.indexes[i].Eval().Float64())
}

func (expr *ArrayAccessExpression) consumeArray(value lib.Value) *lib.ArrayValue {
	switch arr := value.(type) {
	case *lib.ArrayValue:
		return arr
	default:
		panic("Not array value")
	}
}
