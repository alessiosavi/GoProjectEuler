//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package problem9

import (
	"fmt"
	"math"
)

func ExtractCanditate(max float64) [][]float64 {
	var array [][]float64
	var i, j, k float64 = 0, 0, 0
	for i = 0; i < max; i++ {
		for j = 0; j < max; j++ {
			for k = 0; k < max; k++ {
				if i+j+k == 1000 {
					array = append(array, []float64{i, j, k})
				}
			}
		}
	}
	return array
}

func FindPythagoreanTriplet() {
	candiates := ExtractCanditate(1000)
	var exp float64 = 2
	for _, item := range candiates {
		firstTwo := math.Pow(item[0], exp) + math.Pow(item[1], exp)
		if firstTwo == math.Pow(item[2], exp) {
			fmt.Printf("%f\n", (item[0] * item[1] * item[2]))
		}
	}
}
