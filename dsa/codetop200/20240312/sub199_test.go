package _0240312

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{1, 2, 3, 0, 5, 4}, []int{1, 3, 4}},
		{"t2", []int{1, 0, 3}, []int{1, 3}},
		{"t3", []int{1}, []int{1}},
		{"t4", []int{1, 2, 3}, []int{1, 3}},
		{"t5", []int{1, 2, 3, 4}, []int{1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nums)
			got := rightSideView(root)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
