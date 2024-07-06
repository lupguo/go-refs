package _0240314

import (
	. "dsa/data-struct"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 树中路径和为target path
//
func pathSum(root *TreeNode, targetSum int) int {
	// base case
	if root == nil {
		return 0
	}

	// 基于root路径和为targetSum的解个数 + 左、右子树的递归结果
	ans := pathSumFromRoot(root, targetSum)
	return ans + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

// 基于root路径和为targetSum的解个数
func pathSumFromRoot(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var ans int
	if root.Val == targetSum {
		ans++
	}

	// 继续从左右子树根从根节点出发寻求和为remainTargetSum的值
	remainTargetSum := targetSum - root.Val
	return ans + pathSumFromRoot(root.Left, remainTargetSum) + pathSumFromRoot(root.Right, remainTargetSum)
}
