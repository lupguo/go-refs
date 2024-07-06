package _0231208

import (
	"testing"

	. "dsa/data-struct"
)

// https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=leetcode-75
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径定义： 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// 思路: 递归思路，拆分成子问题 - 注意看清楚path定义
//  1. 左、右子树中寻求target解个数 pathSum(root.left, targetSum) + pathSum(root.right, target)
//  2. 基于路径定义，寻求从根节点可能的解: foundPathAns(root, targetSum)
//  3. return pathSum(root.left, targetSum)+pathSum(root.right, targetSum) + foundPathAns(root, targetSum)
func pathSum(root *TreeNode, targetSum int) int {
	// base case
	if root == nil {
		return 0
	}

	// 递归从左、右子树寻求可能的答案
	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + foundPathAns(root, targetSum)
}

func foundPathAns(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var ans int
	if root.Val == targetSum {
		ans++
	}

	// 从子节点找剩余答案
	residueSum := targetSum - root.Val
	return ans + foundPathAns(root.Left, residueSum) + foundPathAns(root.Right, residueSum)
}

func TestPathSum(t *testing.T) {
	tests := []struct {
		name      string
		treeNodes []int
		target    int
		want      int
	}{
		{"t1", []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}, 8, 3},
		{"t2", []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}, 22, 3},
		{"t3", []int{1, 2, 3, 4, 1}, 3, 3},
		{"t4", []int{1}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tTree := IntSliceBFSToBinaryTree(tt.treeNodes)
			if got := pathSum(tTree, tt.target); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
