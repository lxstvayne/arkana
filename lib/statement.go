package lib

import "fmt"

type Statement interface {
	Execute()
	fmt.Stringer
}
