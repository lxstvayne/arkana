package ast

import (
	"strings"
)

type BlockStatement struct {
	statements []Statement
}

func NewBlockStatement(statements []Statement) *BlockStatement {
	return &BlockStatement{
		statements: statements,
	}
}

func (st *BlockStatement) Add(stmt Statement) {
	st.statements = append(st.statements, stmt)
}

func (st *BlockStatement) Execute() {
	for _, stmt := range st.statements {
		stmt.Execute()
	}
}

func (st *BlockStatement) String() string {
	var str strings.Builder

	for _, stmt := range st.statements {
		str.WriteString(stmt.String() + "\n")
	}

	return str.String()
}
