package problem11

import "testing"

func TestInitData(t *testing.T) {
	InitData()
}

func Test(t *testing.T) {
	data := InitData()
	orizzontal := SumOrizzontal(data)
	vertical := SumVertical(data)
	diagonal := SumDiagonal(data)
	inverse := SumDiagonalInverse(data)
	t.Log(orizzontal, vertical, diagonal, inverse)

}

// func TestSumOrizzontal(t *testing.T) {
// 	data := InitData()
// 	t.Log(SumOrizzontal(data))
// }

// func TestSumVertical(t *testing.T) {
// 	data := InitData()
// 	t.Log(SumVertical(data))
// }

// func TestSumDiagonal(t *testing.T) {
// 	data := InitData()
// 	t.Log(SumDiagonal(data))
// }

func TestSumDiagonalInverse(t *testing.T) {
	data := InitData()
	t.Log(SumDiagonalInverse(data))
}
