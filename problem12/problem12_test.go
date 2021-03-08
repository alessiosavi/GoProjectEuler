package problem12

import "testing"

func TestGenerateTriangularNumber(t *testing.T) {
	data, _ := GenerateTriangularNumber(7)

	if data != 1 {
		t.Fail()
	}
}

func TestBruteForceDivisor(t *testing.T) {
	t.Log(BruteForceDivisor(1000))
}

func TestWin(t *testing.T) {
	Win()
}
