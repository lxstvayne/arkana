package ast

type BreakStatement struct {
}

func NewBreakStatement() *BreakStatement {
	return &BreakStatement{}
}

func (st *BreakStatement) Execute() {
	panic("Break Statement")
}

func (st *BreakStatement) String() string {
	return "break"
}
