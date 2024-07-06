package dfs

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
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	vals := []int{root.Val}
	for {
		// 优先看右节点，在看左边的
		if root.Right != nil {
			vals = append(vals, root.Right.Val)
			root = root.Right
			continue
		} else if root.Left != nil {
			vals = append(vals, root.Left.Val)
			root = root.Left
			continue
		}

		break
	}

	return vals
}
