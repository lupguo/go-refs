package _0231220

import (
	. "dsa/data-struct"
)

// 给定一个二叉树 root ，返回其最大深度。
// 思路1: 递归思路, max(maxDepth(root.Left), maxDepth(root.Right))+1
func maxDepthRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// 思路2: 迭代遍历
func maxDepth(root *TreeNode) int {
	var curPath, maxPath int
	// 迭代左子树
	return traverse(root, curPath, maxPath)
}

func traverse(root *TreeNode, curPath int, maxPath int) int {
	if root == nil {
		return maxPath
	}
	curPath++
	maxPath = max(curPath, maxPath)
	leftMaxPath := traverse(root.Left, curPath, maxPath)
	rightMaxPath := traverse(root.Right, curPath, maxPath)
	curPath--

	return max(leftMaxPath, rightMaxPath)
}
