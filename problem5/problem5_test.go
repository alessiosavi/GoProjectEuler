package problem5

import "testing"

func TestCheckSequentialDivisor(t *testing.T) {
	//t.Log(CheckSequentialDivisor(2520, 10))
	var i int = 1
	for {
		if CheckSequentialDivisor(i, 20) {
			t.Log("Divisible ", i)
			break
		}
		i++
	}
}
