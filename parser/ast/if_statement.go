package ast

import (
	"fmt"
)

type IfStatement struct {
	condition Expression
	ifStmt    Statement
	elseStmt  Statement
}

func NewIfStatement(expr Expression, ifStmt Statement, elseStmt Statement) *IfStatement {
	return &IfStatement{
		condition: expr,
		ifStmt:    ifStmt,
		elseStmt:  elseStmt,
	}
}

func (st *IfStatement) Execute() {
	result := st.condition.Eval().Float64()
	if result != 0 {
		st.ifStmt.Execute()
	} else if st.elseStmt != nil {
		st.elseStmt.Execute()
	}
}

func (st *IfStatement) String() string {
	str := fmt.Sprintf("if %s %s ", st.condition, st.ifStmt.String())
	if st.elseStmt != nil {
		str += st.elseStmt.String()
	}
	return str
}
