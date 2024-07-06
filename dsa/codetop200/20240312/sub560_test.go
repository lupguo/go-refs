package _0240312

import (
	"testing"
)

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1}, 0}, 0},
		{"t2", args{[]int{1}, 1}, 1},
		{"t3", args{[]int{1, 2}, 2}, 1},
		{"t4", args{[]int{1, 2}, 3}, 1},
		{"t5", args{[]int{1, 2, 1}, 3}, 2},
		{"t6", args{[]int{1, 2, 1, 0}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
