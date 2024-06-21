package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestTrim(t *testing.T) {
	t.Logf("%#v", strings.Trim(" H  ", ""))
	t.Logf("%#v", strings.Trim(" H  ", " "))
	t.Logf("%#v", strings.Trim(` H \"  `, ` \.`))
	t.Logf("%#v", strings.Trim(`&* dd#?[]{}\|/.  "`, `&*#?[]{}\\|/. "`))
	s := strings.TrimFunc(`1&* 2dd#?3[]{}\|4/.  "`, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
	t.Logf("%#v", s)
}

func TestReplace(t *testing.T) {
	all := strings.ReplaceAll(`a&*#?[]{}\|/. b"`, `&*#?[]{}\\|/. `, "_")
	t.Logf("all => %s", all)
}

func TestStr1VsStr2(t *testing.T) {
	v1 := `10`
	v2 := `9`
	fmt.Print()
	t.Logf("`10` > `9` ? result: %t, %#x, %#x", v1 > v2, v1, v2)
	// t.Logf("`10` > `9` ? result: %t, %x, %x", v1 > v2, v1, v2)
}
