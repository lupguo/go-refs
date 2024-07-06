package _0240311

import (
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2}, 0, 1}, 1},
		{"t2", args{[]int{1, 2, 3}, 0, 2}, 2},
		{"t3", args{[]int{3, 2, 1}, 0, 2}, 0},
		{"t4", args{[]int{2, 3, 1}, 0, 2}, 0},
		{"t5", args{[]int{1, 3, 1}, 0, 2}, 0},
		{"t6", args{[]int{2, 1, 3}, 0, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.nums, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{3, 1, 2}, 1}, 3},
		{"t2", args{[]int{3, 1, 2}, 2}, 2},
		{"t3", args{[]int{3, 1, 2}, 3}, 1},
		{"t4", args{[]int{1, 1, 2}, 3}, 1},
		{"t5", args{[]int{1, 1, 2}, 2}, 1},
		{"t6", args{[]int{1, 1, 2}, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
