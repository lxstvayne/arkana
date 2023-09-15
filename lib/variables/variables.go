package variables

import (
	"arkana/lib"
	"fmt"
	"math"
)

var (
	variables = map[string]lib.Value{
		"PI": lib.NewNumberValue(math.Pi),
		"E":  lib.NewNumberValue(math.E),
	}
)

func IsExists(key string) bool {
	_, ok := variables[key]
	return ok
}

func Get(key string) lib.Value {
	if !IsExists(key) {
		panic(fmt.Sprintf("Const '%s' doesnt exists", key))
	}

	return variables[key]
}

func Set(key string, value lib.Value) {
	variables[key] = value
}
