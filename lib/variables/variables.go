package variables

import (
	"fmt"
	"math"
)

var (
	variables = map[string]float64{
		"PI": math.Pi,
		"E":  math.E,
	}
)

func IsExists(key string) bool {
	_, ok := variables[key]
	return ok
}

func Get(key string) float64 {
	if !IsExists(key) {
		panic(fmt.Sprintf("Const '%s' doesnt exists", key))
	}

	return variables[key]
}

func Set(key string, value float64) {
	variables[key] = value
}
