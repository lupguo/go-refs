package _0231220

import (
	. "dsa/data-struct"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	// 先根 -> 左子树 -> 右子树
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)

	return ret
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	// 左子树 -> 根-> 右子树
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)

	return ret
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	// 左子树 -> 右子树 -> 根
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)

	return ret
}
