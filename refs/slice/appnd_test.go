package slice

import (
	"testing"
)

func TestAppend(t *testing.T) {
	s := []string{"1","2"}
	var k []string
	s = append(s, k...)
	t.Logf("%+v", s)
}
