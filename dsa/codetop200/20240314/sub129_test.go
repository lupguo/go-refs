package _0240314

import (
	"testing"

	. "dsa/data-struct"
)

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1}, 1},
		{"t2", []int{1, 2}, 12},
		{"t3", []int{1, 2, 3}, 25},
		{"t4", []int{1, 2, 0, 3}, 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nums)
			if got := sumNumbers(root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
