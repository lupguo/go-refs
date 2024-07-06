package _0240311

import (
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"a", "b"}, 0},
		{"t2", args{"a", "a"}, 1},
		{"t3", args{"ab", "a"}, 1},
		{"t4", args{"ab", "ab"}, 2},
		{"t5", args{"acb", "ab"}, 2},
		{"t6", args{"acb", "abc"}, 2},
		{"t7", args{"acb", "abcb"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequenceV3(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"a", "bc"}, 0},
		{"t2", args{"ab", "ac"}, 1},
		{"t3", args{"abc", "ac"}, 2},
		{"t4", args{"abc", "acc"}, 2},
		{"t5", args{"abc", "bac"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceV3(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubsequenceV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
