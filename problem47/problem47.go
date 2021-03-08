package problem47

import (
	"fmt"

	mathutils "github.com/alessiosavi/GoGPUtils/math"
)

// The first two consecutive numbers to have two distinct prime factors are:
// 14 = 2 × 7
// 15 = 3 × 5
// The first three consecutive numbers to have three distinct prime factors are:
// 644 = 2² × 7 × 23
// 645 = 3 × 5 × 43
// 646 = 2 × 17 × 19.
// Find the first four consecutive integers to have four distinct prime factors each.
// What is the first of these numbers?

func problem47() {
	findPrimeFactor(20)
}

func findPrimeFactor(n int) []int {
	primes := make([]int, 0)
	// Find primes
	for i := n; i > 0; i-- {
		if mathutils.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Println("Primes: ", primes)

	for k := 0; k < len(primes); k++ {
		for j := k; j < len(primes); j++ {
			// fmt.Printf("%d x %d\n", primes[k], primes[j])
		}
	}

	// mathutils.IsPrime(n)
	return nil
}
