package statements

import "arkana/lib"

type ReturnStatement struct {
	expression lib.Expression
	result     lib.Value
}

func NewReturnStatement(expr lib.Expression) *ReturnStatement {
	return &ReturnStatement{expression: expr}
}

func (st *ReturnStatement) Execute() {
	st.result = st.expression.Eval()
	// RETURN???
}

func (st *ReturnStatement) String() string {
	return "return"
}
