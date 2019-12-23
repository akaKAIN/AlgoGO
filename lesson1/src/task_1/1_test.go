package task_1

import (
	"testing"

)

func TestCalcSum(t *testing.T) {
	result := CalcSum(234)
	if result != 9 {
		t.Error("Error of calculate sum of numbers")
	}
}

