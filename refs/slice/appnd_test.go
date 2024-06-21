package slice

import (
	"testing"
)

func TestAppend(t *testing.T) {
	s := []string{"1", "2"}
	var k []string
	s = append(s, k...)
	t.Logf("%+v", s)
}

func TestAppendNilStructSlice(t *testing.T) {
	type info struct {
		even []int
		odd  []int
	}

	// in := &info{}
	var in *info
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			in.even = append(in.even, i)
		} else {
			in.odd = append(in.odd, i)
		}
	}

	t.Logf("%v", in)
}
