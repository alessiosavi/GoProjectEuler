package problem6

import (
	"testing"
)

func TestCalculateSequentialPow(t *testing.T) {
	if CalculateSequentialPow(10) != 385 {
		t.Fail()
	}
}

func TestCalculateSumPow(t *testing.T) {
	if CalculateSumPow(10) != 3025 {
		t.Fail()
	}
}

func TestTotal(t *testing.T) {
	t.Log(CalculateSumPow(100) - CalculateSequentialPow(100))
}
