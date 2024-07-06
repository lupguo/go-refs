package _0240312

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
//      1           // 1
//  2       3       // 3
// *   5   *   4    // 4
// 思路: 考虑使用BFS，从右向左增加元素
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	que := []*TreeNode{root}
	for len(que) > 0 {
		// 最后一个元素
		last := que[len(que)-1]
		ans = append(ans, last.Val)

		// BFS遍历每层
		var tmpque []*TreeNode
		for _, node := range que {
			if node.Left != nil {
				tmpque = append(tmpque, node.Left)
			}
			if node.Right != nil {
				tmpque = append(tmpque, node.Right)
			}
		}
		que = tmpque
	}

	return ans
}
