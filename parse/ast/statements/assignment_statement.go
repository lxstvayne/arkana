package statements

import (
	"arkana/lib"
	"arkana/lib/variables"
	"fmt"
)

type AssignmentStatement struct {
	variable []rune
	expr     lib.Expression
}

func NewAssignmentStatement(variable []rune, expr lib.Expression) *AssignmentStatement {
	return &AssignmentStatement{
		variable: variable,
		expr:     expr,
	}
}

func (st *AssignmentStatement) Execute() {
	result := st.expr.Eval()
	variables.Set(string(st.variable), result)
}

func (st *AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s", string(st.variable), st.expr)
}
