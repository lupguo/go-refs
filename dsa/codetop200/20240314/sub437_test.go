package _0240314

import (
	"testing"

	. "dsa/data-struct"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		nums      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1}, 1}, 1},
		{"t2", args{[]int{1}, 0}, 0},
		{"t3", args{[]int{1, 2}, 3}, 1},
		{"t4", args{[]int{1, 2, 3}, 3}, 2},
		{"t5", args{[]int{1, 2}, 2}, 1},
		{"t6", args{[]int{1, 2, 3, 4, 5}, 7}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.args.nums)
			if got := pathSum(root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
