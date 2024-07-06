package _0240312

import (
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1}}, 1},
		{"t2", args{[]int{1, 2}}, 2},
		{"t3", args{[]int{3, 1, 2}}, 3},
		{"t4", args{[]int{0, 1}}, 2},
		{"t5", args{[]int{-1, 0, 1}}, 3},
		{"t6", args{[]int{-1, 0, 0, 1}}, 3},
		{"t7", args{[]int{2, -1, 0, 0, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
