package ast

type ContinueStatement struct {
}

func NewContinueStatement() *ContinueStatement {
	return &ContinueStatement{}
}

func (st *ContinueStatement) Execute() {
	panic("Continue Statement")
}

func (st *ContinueStatement) String() string {
	return "continue"
}
