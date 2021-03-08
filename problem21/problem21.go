package problem21

import (
	"log"

	mathutils "github.com/alessiosavi/GoGPUtils/math"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// FindIndexValue is delegated to retrieve the index of the given value into the input array.
// NOTE: FIXME in case of multiple value, only the first will be returned
func FindIndexValue(array []int, value int) int {
	index := -1
	for i := range array {
		if array[i] == value {
			index = i
			break
		}
	}
	return index
}

func RemoveValue(array []int, value int) []int {
	index := FindIndexValue(array, value)
	if index != -1 {
		return remove(array, index)
	}
	return array
}

func FindDivisorsSum(max int) []int {
	var amicable []int
	for i := 0; i < max; i++ {
		// Find divisor for the number i
		divisors1 := mathutils.FindDivisor(i)
		// Remove i itself from the divisor
		divisors1 = RemoveValue(divisors1, i)
		divisor1Sum := mathutils.SumIntArray(divisors1)
		//log.Println("divisor for number1 ", i, " -> ", divisors1, " Sum of divisor -> ", divisor1Sum)
		//log.Println("Sum of divisor1", divisor1Sum)
		// Find divisor for the number divisor1Sum
		divisors2 := mathutils.FindDivisor(divisor1Sum)
		//log.Println("divisor for number ", divisor1Sum, " -> ", divisors2)
		// Remove i itself from the divisor
		divisors2 = RemoveValue(divisors2, divisor1Sum)
		divisor2Sum := mathutils.SumIntArray(divisors2)
		//log.Println("divisor for number2 ", divisor1Sum, " -> ", divisors2, " Sum of divisor -> ", divisor2Sum)
		//log.Println("Sum of divisor1", divisor1Sum)
		addIt := true
		if divisor2Sum == i && divisor1Sum != divisor2Sum {
			for _, alredyPresent := range amicable {
				if divisor1Sum == alredyPresent || divisor2Sum == alredyPresent {
					addIt = false
					break
				}
			}
			if addIt {
				log.Println("Found amicable number: ", divisor1Sum, "-", divisor2Sum)
				amicable = append(amicable, divisor1Sum)
				amicable = append(amicable, divisor2Sum)
			}
		}
	}
	log.Println("List of amicable: ", amicable)
	log.Println("Sum of amicable: ", mathutils.SumIntArray(amicable))
	return amicable
}

func FindEqualValue(array []int) {
	var amicable []int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] == array[j] && i != j {
				log.Println("Find amicable number: ", i, "-", j)
				amicable = append(amicable, i)
				amicable = append(amicable, j)
			}
		}
	}

	log.Println("Sum of amicable: ", mathutils.SumIntArray(amicable))

}
