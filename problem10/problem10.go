package problem10

import (
	"log"
	"math"

	"github.com/alessiosavi/GoGPUtils/helper"
	mathutils "github.com/alessiosavi/GoGPUtils/math"
)

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	mult := float64(2)
	for int(math.Pow(float64(i), mult)) <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func CalculatePrime(max int) {
	array := helper.GenerateSequentialIntArray(max)
	var prime []int
	for i := 0; i < len(array); i++ {
		if IsPrime(array[i]) {
			prime = append(prime, array[i])
		}
	}
	log.Println(mathutils.SumIntArray(prime))
}
