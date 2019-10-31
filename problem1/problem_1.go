//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.`
package problem1

import (
	"log"

	helper "github.com/alessiosavi/GoGPUtils/helper"
	mathutils "github.com/alessiosavi/GoGPUtils/math"
)

var check []int = []int{3, 5}

func GenerateSequentialNumber(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i + 1
	}
	return array
}

func CheckMultiples(n int) bool {
	for i := range check {
		if n%check[i] == 0 {
			return true
		}
	}
	return false
}

func Win() {
	var result []int
	array := helper.GenerateSequentialIntArray(1000)
	for i := range array {
		if CheckMultiples(array[i]) {
			result = append(result, array[i])
		}
	}
	log.Println(mathutils.SumIntArray(result))
}
