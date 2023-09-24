package functions

import (
	"arkana/lib"
	"fmt"
	"math"
)

type SinFunction struct{}

func (sf *SinFunction) Execute(args ...lib.Value) lib.Value {
	if len(args) != 1 {
		panic("One args expected")
	}

	return lib.NewNumberValue(math.Sin(args[0].Float64()))
}

var (
	functions = map[string]lib.Function{
		"sin": &SinFunction{},
	}
)

func IsExists(key string) bool {
	_, ok := functions[key]
	return ok
}

func Get(key string) lib.Function {
	if !IsExists(key) {
		panic(fmt.Sprintf("Function '%s' doesnt exists", key))
	}

	return functions[key]
}

func Set(key string, value lib.Function) {
	functions[key] = value
}
