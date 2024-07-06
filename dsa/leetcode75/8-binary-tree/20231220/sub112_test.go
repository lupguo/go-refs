package _0231220

import (
	"testing"

	. "dsa/data-struct"
)

// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 思路1: 遍历思路，1. 定义一个traver(root)函数, 从根到节点遍历二叉树根节点
//   - 进入节点时候(先序位置），如果存在curSum + root.val == targetSum 返回true
//   - traver(root.Left)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 先序位置，进入节点先加上节点值
	targetSum -= root.Val
	// 到达根节点了，判断curSum是否和targetSum相等
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	// 左子树中是否有满足条件的子节点
	if hasPathSum(root.Left, targetSum) {
		return true
	} else if hasPathSum(root.Right, targetSum) {
		return true
	}

	// 后续位置，从节点离开，扣减掉节点值
	targetSum += root.Val

	// 左右都没有满足条件的值
	return false
}

func TestHasPathTarget(t *testing.T) {
	tests := []struct {
		name   string
		nodes  []int
		target int
		want   bool
	}{
		{"t1", []int{1, 2, 3}, 5, false},
		{"t2", []int{1, 2, 3}, 3, true},
		{"t3", []int{1, 2, 3, 4, 5}, 4, true},
		{"t4", []int{1, 2, 3, 4}, 4, true},
		{"t5", []int{1, 2, 3, 4}, 7, true},
		{"t6", []int{1, 2, 3, 4, 0, 5}, 9, true},
		{"t7", []int{1, 2, 3, 4, 0, 5}, 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nodes)
			if got := hasPathSum(root, tt.target); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
