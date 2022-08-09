package bitx

import (
	"testing"
)

func TestBitMove(t *testing.T) {
	for i := 0; i < 8; i++ {
		t.Logf("i:%032b, i=%[1]d", i)
		t.Logf("n:%032b, i&1=%[1]d", i&1)
		t.Logf("n:%032b, i>>1=%[1]d\n", i>>1)
		t.Log()
	}
}
