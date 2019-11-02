// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package problem3

import (
	"math"
)

// Copied from https://www.geeksforgeeks.org/find-largest-prime-factor-number/
func CalculateMaxPrimeFactor(n int64) int64 {
	var maxPrime int64 = -1
	var i int64
	for n%2 == 0 {
		n /= 2
	}

	for i = 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			n = n / i
		}
	}
	if n > 2 {
		maxPrime = n
	}
	return maxPrime
}
