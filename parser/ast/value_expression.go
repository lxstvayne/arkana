package ast

import (
	"arkana/lib"
)

type ValueExpression struct {
	value lib.Value
}

func NewValueExpression(value interface{}) *ValueExpression {
	var exprValue lib.Value

	switch value := value.(type) {
	case float64:
		exprValue = lib.NewNumberValue(value)
	case string:
		exprValue = lib.NewStringValue(value)
	default:
		panic("Invalid ValueExpression type")
	}

	return &ValueExpression{
		value: exprValue,
	}
}

func (expr *ValueExpression) Eval() lib.Value {
	return expr.value
}

func (expr *ValueExpression) String() string {
	return expr.value.String()
}
