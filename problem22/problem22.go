package problem22

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
	"sort"
)

func Win(filename string) int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bytes.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	names := records[0]
	sort.Strings(names)
	var sum int
	// Subtract 64 due to UPPER CASE ascii character --> 'A' - 64 -> 1
	for i, name := range names {
		var partialSum int
		for _, r := range name {
			partialSum += int(r - 64)
		}
		sum += partialSum * (i + 1)
	}
	return sum
}
