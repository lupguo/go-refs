package slice

import (
	"testing"
)

func TestLenSlice(t *testing.T) {
	var s1 []uint64
	t.Logf("s1:%+v, len(s1)=%d, s1==nil(%t)", s1, len(s1), s1==nil)
}
