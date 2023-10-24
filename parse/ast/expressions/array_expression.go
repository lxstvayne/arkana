package expressions

import (
	"arkana/lib"
	"fmt"

	"github.com/elliotchance/pie/v2"
)

type ArrayExpression struct {
	elements []lib.Expression
}

func NewArrayExpression(elements []lib.Expression) *ArrayExpression {
	return &ArrayExpression{
		elements: elements,
	}
}

func (expr *ArrayExpression) Eval() lib.Value {
	arrayElements := []lib.Value{}
	for _, el := range expr.elements {
		arrayElements = append(arrayElements, el.Eval())
	}

	array := lib.NewArrayValue(arrayElements)

	return array
}

func (expr *ArrayExpression) String() string {
	return fmt.Sprintf("[%s]", pie.Map(expr.elements, func(el lib.Expression) string { return el.String() }))
}
