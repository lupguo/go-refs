package strings

import (
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
