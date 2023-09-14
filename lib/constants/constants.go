package constants

import (
	"fmt"
	"math"
)

var (
	constants = map[string]float64{
		"PI": math.Pi,
		"E":  math.E,
	}
)

func IsExists(key string) bool {
	_, ok := constants[key]
	return ok
}

func Get(key string) float64 {
	if !IsExists(key) {
		panic(fmt.Sprintf("Const '%s' doesnt exists", key))
	}

	return constants[key]
}
