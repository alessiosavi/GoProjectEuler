package problem18

import (
	"log"
	"strconv"
	"strings"

	mathtuils "github.com/alessiosavi/GoGPUtils/math"
	stringutils "github.com/alessiosavi/GoGPUtils/string"
)

const data string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

const testData string = `3
7 4
2 4 6
8 5 9 3`

func RemoveWhiteSpaceString(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := range str {
		if str[i] != 32 {
			b.WriteRune(rune(str[i]))
		}
	}
	return b.String()
}

func LoadBigDataInMatrix(str string) [][]int {
	log.Println("Without space : ", str)
	stringArray := stringutils.Split(str)
	result := make([][]int, len(stringArray))
	for i, array := range stringArray {
		spaceArray := strings.Split(array, " ")
		line := make([]int, len(spaceArray))
		for j, spaceData := range spaceArray {
			tmp, _ := strconv.Atoi(string(spaceData))
			line[j] = tmp
		}
		result[i] = line
	}
	log.Println(result)
	return result
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// FindIndexValue is delegated to retrieve the index of the given value into the input array.
// NOTE: in case of multiple value, only the first will be returned
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

// SortMaxIndex is delegated to return an array that contains the position of the order value (from max to min) of the given array
// {1, 9, 2, 10, 3} -> [3 1 4 2 0] || {7, 6, 5, 4, 3, 2, 1} -> [0 1 2 3 4 5 6] || {1, 2, 3, 4, 5, 6, 7} -> [6 5 4 3 2 1 0]
func SortMaxIndex(array []int) []int {
	var result []int
	var additional int
	var arrayCopy []int = make([]int, len(array))
	copy(arrayCopy, array)
	for i := 0; len(array) > 0; i++ {
		index := mathtuils.MaxIntIndex(array)
		if index != -1 {
			additional = FindIndexValue(arrayCopy, array[index])
			if additional == -1 {
				additional = 0
			}
			result = append(result, additional)
			array = remove(array, index)
		}
	}
	return result
}

func FindPathByIndex(matrix [][]int) {
}

func FindAdiacentPath(matrix [][]int) {
	var result []int
	index := 0
	firstElement := matrix[0][0]
	result = append(result, firstElement)
	for i := 1; i < len(matrix); i++ {
		array := matrix[i]
		if index == 0 {
			// Candidates are the one for index = {0,1}
			if array[0] > array[1] {
				result = append(result, array[0])
			} else {
				result = append(result, array[1])
				index++
			}
		} else {
			// Candidates are the one for index = {index,index+1}
			if array[index] > array[index+1] {
				result = append(result, array[index])
			} else {
				result = append(result, array[index+1])
				index++
			}
		}
	}
	log.Println(result)
	log.Println(mathtuils.SumIntArray(result))
}
