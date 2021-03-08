package problem14

import (
	mathutils "github.com/alessiosavi/GoGPUtils/math"
)

func CollatzProblem(n int) int {
	counter := make([]int, n)
	for i := 1; i < n; i++ {
		k := i
		c := 0
		for k != 1 {
			c++
			if k%2 == 0 {
				k /= 2
			} else {
				k = 3*k + 1
			}
		}
		counter[i] = c
	}
	index := mathutils.MaxIntIndex(counter)
	return index
}
