// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package problem3

import "testing"

func TestCalculatePrimeFactor(t *testing.T) {
	data := CalculateMaxPrimeFactor(600851475143)
	t.Log(data)
}

func BenchmarkCalculatePrimeFactor(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CalculateMaxPrimeFactor(600851475143)
	}
}
