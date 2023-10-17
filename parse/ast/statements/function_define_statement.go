package statements

import (
	"arkana/lib"
	"arkana/lib/functions"
	"fmt"
	"strings"
)

type FunctionDefineStatement struct {
	name     string
	argNames []string
	body     lib.Statement
}

func NewFunctionDefineStatement(name string, args []string, body lib.Statement) *FunctionDefineStatement {
	return &FunctionDefineStatement{
		name:     name,
		argNames: args,
		body:     body,
	}
}

func (fds *FunctionDefineStatement) Execute() {
	functions.Set(fds.name, lib.NewDefineFunction(fds.argNames, fds.body))
}

func (fds *FunctionDefineStatement) String() string {
	return fmt.Sprintf("func (%s) %s", strings.Join(fds.argNames, ","), fds.body.String())
}
