package data_struct

import (
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{1, 2, 3}, []int{1, 2, 3}},
		{"t2", []int{1, 2}, []int{1, 2}},
		{"t3", []int{1, 0, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttTree := IntSliceBFSToBinaryTree(tt.nums)
			bfsGot := ttTree.BFS()
			if !slices.Equal(bfsGot, tt.want) {
				t.Errorf("got %v, but want %v", bfsGot, tt.want)
			}
		})
	}

}
