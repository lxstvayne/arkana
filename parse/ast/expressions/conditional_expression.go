package expressions

import (
	"arkana/lib"
	"fmt"
	"reflect"
	"strings"
)

type ConditionalExpression struct {
	operation Operator
	expr1     lib.Expression
	expr2     lib.Expression
}

func NewConditionalExpression(operation Operator, expr1 lib.Expression, expr2 lib.Expression) *ConditionalExpression {
	return &ConditionalExpression{
		operation: operation,
		expr1:     expr1,
		expr2:     expr2,
	}
}

func (expr *ConditionalExpression) Eval() lib.Value {
	value1 := expr.expr1.Eval()
	value2 := expr.expr2.Eval()

	type1 := reflect.TypeOf(value1)
	type2 := reflect.TypeOf(value2)

	if type1 != type2 {
		panic("Different types")
	}

	var number1, number2 float64

	switch value1.(type) {
	case *lib.StringValue:
		number1 = float64(strings.Compare(value1.String(), value2.String()))
		number2 = 0
	case *lib.NumberValue:
		number1 = value1.Float64()
		number2 = value2.Float64()
	default:
		panic("Not supported expression type")
	}

	var operationResult bool

	switch expr.operation {
	case OPERATOR_LT:
		operationResult = number1 < number2
	case OPERATOR_LTEQ:
		operationResult = number1 <= number2
	case OPERATOR_GT:
		operationResult = number1 > number2
	case OPERATOR_GTEQ:
		operationResult = number1 >= number2
	case OPERATOR_EQUALS:
		operationResult = number1 == number2
	case OPERATOR_NOT_EQUALS:
		operationResult = number1 != number2

	case OPERATOR_AND:
		operationResult = (number1 != 0) && (number2 != 0)
	case OPERATOR_OR:
		operationResult = (number1 != 0) || (number2 != 0)
	default:
		panic("Unsupported Operation")
	}

	var result float64
	if operationResult {
		result = 1
	} else {
		result = 0
	}

	return lib.NewNumberValue(result)
}

func (expr *ConditionalExpression) String() string {
	return fmt.Sprintf("[%s %s %s]", expr.expr1, expr.operation, expr.expr2)
}
