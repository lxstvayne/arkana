package ast

import (
	"fmt"
)

type PrintStatement struct {
	expr Expression
}

func NewPrintStatement(expr Expression) *PrintStatement {
	return &PrintStatement{
		expr: expr,
	}
}

func (st *PrintStatement) Execute() {
	result := st.expr.Eval()
	fmt.Print(result.String())
}

func (st *PrintStatement) String() string {
	return fmt.Sprintf("print %s", st.expr)
}
