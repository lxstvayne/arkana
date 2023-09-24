package ast

type FunctionStatement struct {
	function *FunctionalExpression
}

func NewFunctionStatement(function *FunctionalExpression) *FunctionStatement {
	return &FunctionStatement{function: function}
}

func (st *FunctionStatement) Execute() {
	st.function.Eval()
}

func (st *FunctionStatement) String() string {
	return st.function.String()
}
