package statements

import (
	"arkana/lib"
	"arkana/lib/variables"
	"fmt"
)

type ArrayAssignmentStatement struct {
	variable   string
	idx        lib.Expression
	expression lib.Expression
}

func NewArrayAssignmentStatement(variable string, idx lib.Expression, expr lib.Expression) *ArrayAssignmentStatement {
	return &ArrayAssignmentStatement{
		variable:   variable,
		idx:        idx,
		expression: expr,
	}
}

func (st *ArrayAssignmentStatement) Execute() {
	variable := variables.Get(st.variable)
	switch array := variable.(type) {
	case *lib.ArrayValue:
		array.Set(int(st.expression.Eval().Float64()), st.expression.Eval())
	default:
		panic("Array expected!")
	}
}

func (st *ArrayAssignmentStatement) String() string {
	return fmt.Sprintf("%s [%s] = %s", st.variable, st.idx, st.expression)
}
