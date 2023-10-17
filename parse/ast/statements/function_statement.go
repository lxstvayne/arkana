package statements

import "arkana/parse/ast/expressions"

type FunctionStatement struct {
	function *expressions.FunctionalExpression
}

func NewFunctionStatement(function *expressions.FunctionalExpression) *FunctionStatement {
	return &FunctionStatement{function: function}
}

func (st *FunctionStatement) Execute() {
	st.function.Eval()
}

func (st *FunctionStatement) String() string {
	return st.function.String()
}
