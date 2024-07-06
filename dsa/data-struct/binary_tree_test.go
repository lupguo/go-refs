package data_struct

import (
	"testing"
)

func TestIntSliceBFSToBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"t1", []int{1, 2, 3, 0, 4, 5}},
		{"t2", []int{1, 2, 3, 0, 4, 5, 6, 0, 0, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", IntSliceBFSToBinaryTree(tt.vals))
		})
	}
}
