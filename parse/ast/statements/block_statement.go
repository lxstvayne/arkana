package statements

import (
	"arkana/lib"
	"strings"
)

type BlockStatement struct {
	statements []lib.Statement
}

func NewBlockStatement(statements []lib.Statement) *BlockStatement {
	return &BlockStatement{
		statements: statements,
	}
}

func (st *BlockStatement) Add(stmt lib.Statement) {
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
