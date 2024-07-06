package _0240312

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"t1", []int{1, 2, 3, 0, 5, 6, 7}, [][]int{{1}, {3, 2}, {5, 6, 7}}},
		{"t2", []int{1, 2, 3, 0, 0, 6, 7}, [][]int{{1}, {3, 2}, {6, 7}}},
		{"t3", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nums)
			got := zigzagLevelOrder(root)
			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
			for i, levels := range got {
				if !slices.Equal(tt.want[i], levels) {
					t.Errorf("got %v, but want %v", got, tt.want)
					break
				}
			}
		})
	}
}
