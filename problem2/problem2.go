//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package problem2

import (
	"log"
)

func ExtractEvenValuedNumber(array []int64) []int64 {
	var result []int64
	for i := range array {
		if array[i]%2 == 0 {
			result = append(result, array[i])
		}
	}
	return result
}

func GenerateFibonacci(max int64) []int64 {
	var array []int64
	// Hardcoded for enhance for performance
	array = append(array, 1)
	array = append(array, 2)
	i := 2
	var value int64 = array[i-1] + array[i-2]
	for value < max {
		value = array[i-1] + array[i-2]
		i++
		array = append(array, value)
	}
	log.Println(array)
	return array
}
