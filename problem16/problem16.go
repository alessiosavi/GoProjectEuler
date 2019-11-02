package problem16

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Problem16(n float64) {
	number := math.Pow(2, n)
	log.Println(number)
	m := fmt.Sprintf("%f", number)
	sum := 0
	for i := range m {
		data, _ := strconv.Atoi(string(m[i]))
		sum += data
	}
	log.Println(sum)
}
