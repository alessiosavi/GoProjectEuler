// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible ( -> with no reminder) by all of the numbers from 1 to 20?
package problem5

import "log"

func CheckSequentialDivisor(n, target int) bool {
	for i := target; i > 1; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func Win() {
	var i int = 1
	for {
		if CheckSequentialDivisor(i, 20) {
			log.Println(i)
			break
		}
		i++
	}
}
