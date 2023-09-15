package lib

import "strconv"

type StringValue struct {
	value string
}

func NewStringValue(value string) *StringValue {
	return &StringValue{
		value: value,
	}
}

func (value *StringValue) Float64() float64 {
	f, err := strconv.ParseFloat(value.value, 64)
	if err != nil {
		return 0
	}
	return f
}

func (value *StringValue) String() string {
	return value.value
}
