package _0230901

import (
	"reflect"
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{"t1", "AB", "ABAB", "AB"},
		{"t2", "ABC", "ABAB", ""},
		{"t3", "ABAB", "ABABABAB", "ABAB"},
		{"t4", "AB", "ABABABAB", "AB"},
		{
			"t5",
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{"t6", "LEET", "ROOT", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("case[%s] got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_maxLenGcd(t *testing.T) {
	tests := []struct {
		name  string
		gcds1 []string
		gcds2 []string
		want  string
	}{
		{"t1", []string{"AB"}, []string{"AB"}, "AB"},
		{"t2", []string{"AB", "ABAB"}, []string{"AB"}, "AB"},
		{"t3", []string{"AB", "ABAB"}, []string{"ABC"}, ""},
		{"t4", []string{"AB", "ABAB"}, []string{"AB", "ABAB"}, "ABAB"},
		{"t5", []string{"AA", "AAA", "AAAAA"}, []string{"AA", "AAAAA"}, "AAAAA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLenGcd(tt.gcds1, tt.gcds2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("case[%s] got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_gcdListOfString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{"ABAB"}, []string{"AB"}},
		{"t2", args{"ABA"}, []string{"ABA"}},
		{"t3", args{"ABABABAB"}, []string{"AB", "ABAB"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdListOfString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gcdListOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
