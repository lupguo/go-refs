package _02404

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"t1", "", 0},
		{"t2", "a", 1},
		{"t3", "aa", 1},
		{"t4", "aab", 2},
		{"t5", "aaba", 2},
		{"t6", "ababa", 2},
		{"t7", "abacba", 3},
		{"t8", "dvdf", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkWinOk(t *testing.T) {
	type args struct {
		b      byte
		window []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{'a', []byte{'a'}}, true},
		{"t2", args{'a', []byte{'b'}}, false},
		{"t3", args{'a', []byte{'a', 'b'}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winCheckOk(tt.args.b, tt.args.window); got != tt.want {
				t.Errorf("winCheckOk() = %v, want %v", got, tt.want)
			}
		})
	}
}
