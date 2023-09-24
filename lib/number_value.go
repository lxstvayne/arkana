package lib

import (
	"fmt"
)

type NumberValue struct {
	value float64
}

func NewNumberValue(value float64) *NumberValue {
	return &NumberValue{
		value: value,
	}
}

func (value *NumberValue) Float64() float64 {
	return value.value
}

func (value *NumberValue) String() string {
	s := fmt.Sprintf("%v", value.value)
	return s
}
