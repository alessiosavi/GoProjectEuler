package problem4

import (
	"log"
	"strconv"
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	t.Log(CheckPalindrome("abba"))
	t.Log(CheckPalindrome("abcba"))
}
func TestGenerateThreeDigitArray(t *testing.T) {
	k := 0
	for i := 99999999; i > 90000000; i-- {
		for j := 99999999; j > 90000000; j-- {
			k++
			n := strconv.Itoa(i * j)
			if CheckPalindrome(n) {
				log.Println("Palindrome -> ", n, " Itearazioni: ", k)
				return
			}
		}
	}
}

func TestWin(t *testing.T) {
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
