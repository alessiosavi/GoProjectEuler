package problem13

import "testing"

func TestPadArray(t *testing.T) {
	array := []int{1, 2, 3, 4}
	t.Log(PadArray(array, 5))
}

func TestSumArrays(t *testing.T) {
	array1 := []int{1, 1, 2, 3, 4}
	array2 := []int{9, 3, 3, 3}
	// 10567
	t.Log(SumArrays(array1, array2))
}

func TestReverseArray(t *testing.T) {
	array2 := []int{1, 2, 3, 4}
	t.Log(ReverseArray(array2))
}

func TestSumBig(t *testing.T) {
	result := SumBig()
	res := result[0:10]
	t.Log("Result :", JoinInts(res))
}
