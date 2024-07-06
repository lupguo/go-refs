package basex

import (
	"testing"
)

func TestDivision(t *testing.T) {
	v1 := 3.0 / 10.0
	v2 := 3.0 / 10
	v3 := 3 / 10.0

	// print 0.3
	n := 3
	v4 := float64(n) / 10

	// print 0
	v5 := n / 10.0

	t.Logf("%+v, %+v, %+v, %+v, %+v", v1, v2, v3, v4, v5)
}
