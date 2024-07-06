package _0231225

import (
	. "dsa/data-struct"
)

// 思路: 基于原本递归思路
func flatten(root *TreeNode) {
	// base条件
	if root == nil {
		return
	}

	// 左子树转成链表
	flatten(root.Left)
	// 右子树转成链表
	flatten(root.Right)

	// 左侧处理
	left := root.Left
	root.Left = nil
	if left == nil {
		return
	}

	// 右侧(遍历左链表，尾部加上右链表)
	p := left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root.Right

	// 接上右链表
	root.Right = left
}
