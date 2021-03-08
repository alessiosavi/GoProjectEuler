package problem23

import "testing"

const max int = 28123

func TestCheckPerfectNumber(t *testing.T) {
	t.Log(CheckPerfectNumber(28))
}

func TestFindAbundantNumber(t *testing.T) {
	n := FindAbundantNumber(max)
	t.Log("From 0 to ", max, " there are ", len(n), " abundant number")
}

func TestWin(t *testing.T) {
	Win(max)
}
