package tregex

import (
	"regexp"
	"testing"
)

func TestPhone11(t *testing.T) {
	var phoneReg = regexp.MustCompile(`^\d{11}$`)
	b := phoneReg.MatchString("1234567890123111")
	t.Logf("b=%t", b)
}
