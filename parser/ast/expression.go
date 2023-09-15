package ast

import (
	"arkana/lib"
	"fmt"
)

type Expression interface {
	Eval() lib.Value
	fmt.Stringer
}
