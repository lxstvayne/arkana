package expressions

import (
	"arkana/lib"
)

var (
	NONE = &NoneExpression{}
)

type NoneExpression struct{}

func (expr *NoneExpression) Eval() lib.Value {
	return lib.NewNumberValue(0)
}

func (expr *NoneExpression) String() string {
	return "None"
}
