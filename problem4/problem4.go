// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
package problem4

import (
	"log"
	"strconv"
)

func GenerateThreeDigitArray() []int {
	var max int = 999
	array := make([]int, max-99)
	var i int = 0
	for max > 99 {
		array[i] = max
		i++
		max--
	}
	return array
}

func CheckPalindrome(str string) bool {
	length := len(str) - 1
	for i := range str {
		if str[i] != str[length-i] {
			//log.Println("", str[i], " is different from ", str[length-i])
			return false
		}
	}
	return true
}
func Win() {
	k := 0
	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			k++
			n := strconv.Itoa(i * j)
			if CheckPalindrome(n) {
				log.Println(n)
				return
			}
		}
	}
}
