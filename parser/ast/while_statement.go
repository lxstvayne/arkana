package ast

import (
	"fmt"
)

type WhileStatement struct {
	condition Expression
	statement Statement
}

func NewWhileStatement(condition Expression, statement Statement) *WhileStatement {
	return &WhileStatement{
		condition: condition,
		statement: statement,
	}
}

func (st *WhileStatement) Execute() {
	for {
		if st.condition.Eval().Float64() == 0 {
			break
		}

		st.statement.Execute()
	}
}

func (st *WhileStatement) String() string {
	return fmt.Sprintf("while %s %s", st.condition.String(), st.statement.String())
}
