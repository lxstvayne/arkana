package ast

import "fmt"

type Expression interface {
	Eval() float64
	fmt.Stringer
}
