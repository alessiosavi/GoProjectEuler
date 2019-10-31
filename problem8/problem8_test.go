package problem8

import "testing"

func TestSplitStringInArray(t *testing.T) {
	temp := SplitStringInArray(data)
	MultiplyNearestNumber(temp, 13)
}

func TestMultiplyArray(t *testing.T) {
	array := []int{9, 9, 8, 9}
	t.Log(MultiplyArray(array))
}
