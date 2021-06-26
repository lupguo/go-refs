package structs

import (
	"testing"
)

type K struct {
	Num *int
}

func TestDefaultValue(t *testing.T) {
	k1 := new(K)
	n := 0
	k2 := &K{Num: &n}
	t.Logf("%#v", k1)
	t.Logf("%#v", k2)
}
