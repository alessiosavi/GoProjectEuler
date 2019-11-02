package problem22

import (
	"testing"
)

const filename string = `/tmp/p022_names.txt`

func TestLoadCsv(t *testing.T) {
	t.Log(Win(filename))
}
