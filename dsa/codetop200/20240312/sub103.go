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
// 即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
func zigzagLevelOrder(root *TreeNode) [][]int {
	return bfsTraverse(root, true)
}

// 层次遍历，每层根据isLeft做处理，决定遍历的方向
func bfsTraverse(root *TreeNode, isLeft bool) [][]int {
	// base case
	if root == nil {
		return nil
	}

	que := []*TreeNode{root}
	var ret [][]int
	for len(que) > 0 {
		//    1 -- 1 [left]
		//  2  3 -- 3 -> 2 [right]
		// * 5 6 7 -- 5, 6,7 [left]
		var levels []int
		var tmpque []*TreeNode
		for _, node := range que {
			if node == nil {
				continue
			}
			if isLeft {
				levels = append(levels, node.Val)
			} else {
				levels = append([]int{node.Val}, levels...)
			}

			if node.Left != nil {
				tmpque = append(tmpque, node.Left)
			}
			if node.Right != nil {
				tmpque = append(tmpque, node.Right)
			}
		}

		// 层级元素添加
		ret = append(ret, levels)

		// 方向变换
		isLeft = !isLeft

		// queue更新
		que = tmpque
	}
	return ret
}
