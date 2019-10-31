package problem12

import (
	"log"
	"math"
)

func GenerateTriangularNumber(n int) (int, int) {
	return (n * (n + 1)) / 2, n
}

func BruteForceDivisor(n int) int {
	var count int = 0
	max := int(math.Sqrt(float64(n)))
	for i := 1; i <= max; i++ {
		if n%i == 0 {
			div := n / i
			if div != i {
				count += 2
			} else {
				count++
			}
		}
	}
	return count
}

func Win() {
	i := 1
	for {
		n, index := GenerateTriangularNumber(i)
		divisor := BruteForceDivisor(n)
		log.Println(divisor)
		if divisor > 500 {
			log.Println("Number: ", n, " Index: ", index)
			break
		}
		i++
	}
}
