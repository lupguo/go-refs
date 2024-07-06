package week17

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{0}, 1},
		{"t2", []int{1}, 1},
		{"t3", []int{0, 1}, 2},
		{"t4", []int{0, 0}, 2},
		{"t5", []int{0, 0, 1}, 3},
		{"t6", []int{0, 0, 2}, 2},
		{"t7", []int{0, 0, 2, 3, 4}, 3},
		{"t8", []int{0, 0, 2, 2, 4}, 2},
		{"t9", []int{0, 0, 2, 2, 3, 4}, 4},
		{"t10", []int{0, 0, 1, 2, 2, 3, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
