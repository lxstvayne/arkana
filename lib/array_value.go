package lib

import (
	"fmt"
	"strings"

	"github.com/elliotchance/pie/v2"
)

type ArrayValue struct {
	elements []Value
}

func NewArrayValue(elements []Value) *ArrayValue {
	return &ArrayValue{
		elements: elements,
	}
}

func (value *ArrayValue) Copy() *ArrayValue {
	return NewArrayValue(value.elements)
}

func (value *ArrayValue) Get(idx int) Value {
	if idx < 0 {
		return value.elements[len(value.elements)+idx]
	}
	return value.elements[idx]
}

func (value *ArrayValue) Set(idx int, v Value) {
	value.elements[idx] = v
}

func (value *ArrayValue) Float64() float64 {
	panic("Cannot cast array to float64")
}

func (value *ArrayValue) String() string {
	return fmt.Sprintf("[%s]", strings.Join(pie.Map(value.elements, func(v Value) string { return v.String() }), ", "))
}
