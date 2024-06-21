package leetcode

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"11", "10"}, "101"},
		{"t2", args{"0", "0"}, "0"},
		{"t3", args{"0", "1"}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary2(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"11", "10"}, "101"},
		{"t2", args{"0", "0"}, "0"},
		{"t3", args{"0", "1"}, "1"},
		{"t4", args{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"},
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary2() = %v, want %v", got, tt.want)
			}
		})
	}
}
