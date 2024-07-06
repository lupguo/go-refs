package basex

import (
	"math"
	"testing"
)

func TestFloat(t *testing.T) {
	f1, f2 := 3.14, 3.14
	f3, f4 := math.Pi, math.Pi
	t.Logf("%T, %[1]T", f1)
	t.Logf("f1==f2 got %v", f1 == f2)
	t.Logf("f3==f4 got %v", f3 == f4)

	f5 := float32(1) / float32(3)
	f6 := float32(0.33333333333333333333)
	f7 := float32(0.333333)
	t.Logf("f5==f6 got %v, f5=%v, f6=%v", f5 == f6, f5, f6)
	t.Logf("f5==f7 got %v, f5=%v, f7=%v", f5 == f7, f5, f7)
}
