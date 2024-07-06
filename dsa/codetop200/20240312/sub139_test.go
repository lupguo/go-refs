package _0240312

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"a", []string{"a"}}, true},
		{"t2", args{"a", []string{"b"}}, false},
		{"t3", args{"a", []string{"b", "a"}}, true},
		{"t4", args{"ab", []string{"b", "a"}}, true},
		{"t5", args{"ba", []string{"b", "a"}}, true},
		{"t6", args{"baa", []string{"b", "a"}}, true},
		{"t7", args{"abb", []string{"b", "a"}}, true},
		{"t8", args{"abbc", []string{"b", "a"}}, false},
		{"t9", args{"abb", []string{"ab", "a"}}, false},
		{"t10", args{"aba", []string{"ab", "a"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
