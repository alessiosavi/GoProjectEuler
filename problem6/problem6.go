//The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package problem6

import (
	"log"
	"math"
)

func CalculateSequentialPow(max int) int64 {
	var total int64 = 0
	for i := 0; i <= max; i++ {
		iF := float64(i)
		total += int64(math.Pow(iF, 2))
	}
	log.Println("sum of the squares for ", max, " is ", total)
	return total
}

func CalculateSumPow(max int) int64 {
	var total int = 0
	for i := 0; i <= max; i++ {
		total += i
	}
	data := math.Pow(float64(total), 2)
	log.Println(data)
	log.Println("square of the sum ", max, " is ", total)
	return int64(data)
}
