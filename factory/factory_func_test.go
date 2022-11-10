package factory

import (
	"testing"
)

func TestOperator(t *testing.T) {
	var pof OperatorFactory
	var opr MathOperator
	pof = &PlusOperatorFactory{}
	opr = pof.Create()
	opr.Compute()
}
