package _0231220

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// 二叉树反转
//
//	思路1: 递归的反转子树，后序遍历位置处理根节点左、右子树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left

	return root
}

// 思路2: 迭代方式，先序位置，针对root节点，交换左、右子树
func invertTreeRange(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  []int
	}{
		{"t1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"t2", []int{1, 2}, []int{1, 2}},
		{"t3", []int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
	}
	for _, tt := range tests {
		root := IntSliceBFSToBinaryTree(tt.nodes)
		invertRoot := invertTreeRange(root)
		if got := invertRoot.BFS(); !slices.Equal(got, tt.want) {
			t.Errorf("got %v, but want %v", got, tt.want)
		}
	}
}
