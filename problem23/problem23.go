package problem23

import (
	"log"

	mathutils "github.com/alessiosavi/GoGPUtils/math"
	"github.com/alessiosavi/ProjectEulerGo/problem21"
)

func CheckPerfectNumber(n int) bool {
	divisors1 := mathutils.FindDivisor(n)
	// Remove i itself from the divisor
	divisors1 = problem21.RemoveValue(divisors1, n)
	divisor1Sum := mathutils.SumIntArray(divisors1)
	return divisor1Sum == n
}

func CheckAbundantNumber(n int) bool {
	divisors1 := mathutils.FindDivisor(n)
	// Remove i itself from the divisor
	divisors1 = problem21.RemoveValue(divisors1, n)
	divisor1Sum := mathutils.SumIntArray(divisors1)
	return divisor1Sum > n
}

func CheckDeficentNumber(n int) bool {
	divisors1 := mathutils.FindDivisor(n)
	// Remove i itself from the divisor
	divisors1 = problem21.RemoveValue(divisors1, n)
	divisor1Sum := mathutils.SumIntArray(divisors1)
	return divisor1Sum < n
}

func FindDeficentNumber(max int) []int {
	var deficentNumber []int
	for i := 1; i <= max; i++ {
		if CheckDeficentNumber(i) {
			deficentNumber = append(deficentNumber, i)
		}
	}
	return deficentNumber
}

func FindAbundantNumber(max int) []int {
	var abundantNumber []int
	for i := 1; i <= max; i++ {
		if CheckAbundantNumber(i) {
			// log.Println("Number ", i, " is an abundant number")
			abundantNumber = append(abundantNumber, i)
		}
	}
	return abundantNumber
}

func Win(max int) int {
	var abundantNumber []int = FindAbundantNumber(max)
	var canWriteAsAbundant map[int]struct{} = make(map[int]struct{})
	for i := 0; i < len(abundantNumber); i++ {
		for j := i; j < len(abundantNumber); j++ {
			canWriteAsAbundant[abundantNumber[i]+abundantNumber[j]] = struct{}{}
		}
	}

	var continuosInt map[int]struct{} = make(map[int]struct{})

	for i := 1; i < max; i++ {
		continuosInt[i] = struct{}{}
	}
	for key := range canWriteAsAbundant {
		delete(continuosInt, key)
	}
	var sum int
	for key := range continuosInt {
		sum += key
	}
	log.Println("Sum: ", sum)
	return sum
}
