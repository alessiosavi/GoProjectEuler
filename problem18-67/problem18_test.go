package problem18

import (
	"log"
	"testing"
)

func TestSortMaxIndex(t *testing.T) {
	var array []int = []int{1, 2, 3, 4, 5, 6, 7}
	// var array []int = []int{7, 6, 5, 4, 3, 2, 1}
	//var array []int = []int{1, 9, 2, 10, 3}

	log.Println(SortMaxIndex(array))
}

func TestSortMaxIndexBig(t *testing.T) {
	//var array []int = []int{1, 2, 3, 4, 5, 6, 7}
	// var array []int = []int{7, 6, 5, 4, 3, 2, 1}
	//var array []int = []int{1, 9, 2, 10, 3}
	matrix := LoadBigDataInMatrix(data)
	for i := range matrix {
		log.Println(SortMaxIndex(matrix[i]))
	}
}

func TestLoadBigDataInMatrix(t *testing.T) {
	LoadBigDataInMatrix(data)
}

func TestFindPath(t *testing.T) {
	FindAdiacentPath(LoadBigDataInMatrix(testData))
}
func TestBigFindPath(t *testing.T) {
	FindAdiacentPath(LoadBigDataInMatrix(data))
}

func TestFindPathByIndex(t *testing.T) {
	FindPathByIndex(LoadBigDataInMatrix(data))
}

func TestProblem67(t *testing.T) {
	matrix := LoadBigDataInMatrix(problem67)
	t.Log(BacktrackMatrix(matrix))
}
func BenchmarkProblem67(t *testing.B) {
	matrix := LoadBigDataInMatrix(problem67)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		BacktrackMatrix(matrix)
	}
}

func TestProblem18(t *testing.T) {
	matrix := LoadBigDataInMatrix(data)
	BacktrackMatrix(matrix)
}

/*
Matrix of index related to highest value
[0]
[0 1]
[2 1 0]
[2 1 0 3]
[2 4 3 0 1]
[3 5 2 0 4 1]
[0 2 3 6 5 4 1]
[0 7 6 1 3 5 4 2]
[4 6 7 3 0 0 5 8 2]
[8 2 1 4 0 6 3 5 9 7]
[8 6 1 3 0 7 9 2 5 4 10]
[7 4 5 0 9 11 8 2 3 6 6 1]
[0 0 1 8 2 9 12 7 3 11 10 4 5]
[4 11 8 10 3 6 1 0 5 12 13 7 9 2]
[2 2 9 8 6 1 12 11 10 3 4 4 5 0 0]

75
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
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

*/
