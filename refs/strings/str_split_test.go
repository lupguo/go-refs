package strings

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{"t1", ""},
		{"t2", "1,2"},
		{"t3", "1,1,2"},
		{"t4", "1,1,2,"},
		{"t5", ",,2,"},
		{"t6", "a ,,, "},
		{"t7", ", "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trimStr := strings.Trim(tt.str, ", ")
			t.Logf("strings.Split('%s', ',')=%#v, len=%d\n", tt.str, strings.Split(trimStr, ","), len(strings.Split(trimStr, ",")))
		})
	}
	t.Logf("%#v, len=%d", []string{}, len([]string{}))
}

func TestSplitCase2(t *testing.T) {
	for _, v := range strings.Split("        ", " ") {
		t.Logf("`%+v`, len(v)=%d", v, len(v))
	}

}
