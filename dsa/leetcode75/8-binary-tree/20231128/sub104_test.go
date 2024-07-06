package _0231128

import (
	"testing"

	. "dsa/data-struct"
)

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// 思路:
//  1. 申请两个变量 depth, maxDepth
//  2. 从根节点开始向下进行DFS深度遍历(递归)，如果存在左节点，depth++，若超过maxDepth，则更新maxDepth，同理，计算右节点的最大深度
//  3. 最后返回，左右子树的最大节点
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return GetTreeDepth(root, 1)
}

func GetTreeDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	// 左节点非空
	leftDepth, rightDepth := depth, depth
	if node.Left != nil {
		leftDepth++
		leftDepth = GetTreeDepth(node.Left, leftDepth)
	}

	// 右节点非空
	if node.Right != nil {
		rightDepth++
		rightDepth = GetTreeDepth(node.Right, rightDepth)
	}

	return max(leftDepth, rightDepth)
}

func TestGetTreeDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"t1", IntSliceBFSToBinaryTree([]int{1, 0, 2}), 2},
		{"t2", IntSliceBFSToBinaryTree([]int{3, 9, 20, 0, 0, 15, 7}), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTreeDepth(tt.root, 1); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
