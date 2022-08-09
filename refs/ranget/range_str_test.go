package ranget

import (
	"testing"
)

func TestRangeString(t *testing.T) {
	s := "abcd"
	for i, v := range s {
		t.Logf("i=%#v, i's type is %T, v=%#v", i, i, v)
	}

	s = `abcd`
	for i, v := range s {
		t.Logf("i=%#v, i's type is %T, v=%#v", i, i, v)
	}

	for v := range s {
		t.Logf("v=%v", v)
	}
}

func TestRangeRune(t *testing.T) {

}
