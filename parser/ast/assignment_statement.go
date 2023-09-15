package ast

import (
	"arkana/lib/variables"
	"fmt"
)

type AssignmentStatement struct {
	variable []rune
	expr     Expression
}

func NewAssignmentStatement(variable []rune, expr Expression) *AssignmentStatement {
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
