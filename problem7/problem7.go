package problem7

import "math"

func IsPrimitive(n int) bool {
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
