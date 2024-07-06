package _0240311

import (
	"testing"
	"unsafe"
)

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"a", ""}, 'a'},
		{"t2", args{"ab", "b"}, 'a'},
		{"t3", args{"ab", "bc"}, 'a' + 'c'},
		{"t4", args{"abc", "bc"}, 'a'},
		{"t5", args{"abc", "bac"}, 'a' + 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumAscii(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"a"}, 97},
		{"t2", args{"z"}, 97 + 26 - 1},
		{"t3", args{"0"}, 48},
		{"t3", args{"9"}, '9'},
		{"t4", args{"a"}, 'a'},
		{"t5", args{"A"}, 65},
		{"t6", args{"Z"}, 65 + 26 - 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAscii(tt.args.s); got != tt.want {
				t.Errorf("sumAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAscii(t *testing.T) {
	t.Log('a')
	var m int
	var n int32
	t.Logf("int m size=%v, int n size=%v", unsafe.Sizeof(m), unsafe.Sizeof(n))
}
