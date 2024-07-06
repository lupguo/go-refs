package _0240314

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1}, 1},
		{"t2", []int{1, 2}, 3},
		{"t3", []int{1, 2, 3}, 6},
		{"t4", []int{1, 2, -1}, 3},
		{"t5", []int{1, 2, -1, 2}, 4},
		{"t6", []int{1, -2, 1, 2}, 3},
		{"t7", []int{2, -1, 1, 2}, 4},
		{"t8", []int{2, -1, 3, -2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
