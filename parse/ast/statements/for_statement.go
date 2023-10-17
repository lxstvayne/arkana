package statements

import (
	"arkana/lib"
	"fmt"
)

type ForStatement struct {
	initialization lib.Statement
	termination    lib.Expression
	increment      lib.Statement
	block          lib.Statement
}

func NewForStatement(initialization lib.Statement, termination lib.Expression, increment lib.Statement, block lib.Statement) *ForStatement {
	return &ForStatement{
		initialization: initialization,
		termination:    termination,
		increment:      increment,
		block:          block,
	}
}

func (st *ForStatement) Execute() {
	st.initialization.Execute()
	for {
		if st.termination.Eval().Float64() == 0 {
			break
		}

		st.block.Execute()
		st.increment.Execute()
	}
}

func (st *ForStatement) String() string {
	return fmt.Sprintf("for %s, %s, %s, %s", st.initialization.String(), st.termination.String(), st.increment.String(), st.block.String())
}
