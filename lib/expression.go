package lib

import (
	"fmt"
)

type Expression interface {
	Eval() Value
	fmt.Stringer
}
