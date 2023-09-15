package ast

import "fmt"

type Statement interface {
	Execute()
	fmt.Stringer
}
