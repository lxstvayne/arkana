package ast

import (
	"arkana/lib"
	"fmt"
	"strings"
)

type BinaryExpression struct {
	operation rune
	expr1     Expression
	expr2     Expression
}

func NewBinaryExpression(operation rune, expr1 Expression, expr2 Expression) *BinaryExpression {
	return &BinaryExpression{
		operation: operation,
		expr1:     expr1,
		expr2:     expr2,
	}
}

func (expr *BinaryExpression) Eval() lib.Value {
	value1 := expr.expr1.Eval()
	value2 := expr.expr2.Eval()

	switch value1.(type) {
	case *lib.StringValue:
		switch expr.operation {
		case rune('+'):
			return lib.NewStringValue(value1.String() + value2.String())
		case rune('*'):
			// FIXME: TYPE CHECKING
			iterations := int(value2.Float64())
			return lib.NewStringValue(strings.Repeat(value1.String(), iterations))
		default:
			panic("Invalid operation")
		}
	}

	number1 := value1.Float64()
	number2 := value2.Float64()
	switch expr.operation {
	case rune('+'):
		return lib.NewNumberValue(number1 + number2)
	case rune('-'):
		return lib.NewNumberValue(number1 - number2)
	case rune('*'):
		return lib.NewNumberValue(number1 * number2)
	case rune('/'):
		return lib.NewNumberValue(number1 / number2)
	default:
		// FIXME
		return lib.NewNumberValue(number1 + number2)
	}
}

func (expr *BinaryExpression) String() string {
	return fmt.Sprintf("[%s %c %s]", expr.expr1, expr.operation, expr.expr2)
}
