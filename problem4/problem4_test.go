package problem4

import (
	"sort"
	"strconv"
	"testing"
)

func TestGenerateThreeDigitArray(t *testing.T) {
	data1 := GenerateThreeDigitArray()
	var data2 []int = make([]int, len(data1))
	copy(data2, data1)
	data2[0] = -999
	t.Log(data1)
	t.Log(data2)
	var result []int
	for i := range data1 {
		for j := range data2 {
			result = append(result, data1[i]*data2[j])
		}
	}
	sort.Ints(result)
	var stringArray []string = make([]string, len(result))
	for i := range result {
		stringArray[i] = strconv.Itoa(result[i])
	}
	t.Log(cap(result))
	t.Log(cap(stringArray))
	for i := len(stringArray) - 1; i > -1; i-- {
		if CheckPalindrome(stringArray[i]) {
			t.Log(stringArray[i])
			return
		}
	}
}

func TestCheckPalindrome(t *testing.T) {
	t.Log(CheckPalindrome("abba"))
	t.Log(CheckPalindrome("abcba"))

}
