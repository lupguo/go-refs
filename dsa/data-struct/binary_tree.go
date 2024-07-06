package data_struct

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BFS 遍历的打印一颗树
func (root *TreeNode) String() string {
	if root == nil {
		return "[]"
	}
	queue := []*TreeNode{root}

	// BFS遍历队列
	var vals []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// BFS先打印当前层
		vals = append(vals, node.Val)

		// 当前节点的左、右节点存在非空，则加入到queue中
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	ret := "["
	for i, val := range vals {
		if i != len(vals)-1 {
			ret += fmt.Sprintf("%d,", val)
		} else {
			ret += fmt.Sprintf("%d", val)
		}

	}
	ret += "]"

	return ret
}

// IntSliceBFSToBinaryTree 整型Slice转二叉树制树
// root = [3,9,20,null,null,15,7]
//
//	  3
//	 / \
//	9   20
//	    / \
//	   15  7
func IntSliceBFSToBinaryTree(vals []int) *TreeNode {
	// return buildCompleteBinaryTreeByRec(vals, 0)
	return buildBinaryTreeByLoop(vals)
}

// 递归(类似深度遍历方式构建节点)
func buildCompleteBinaryTreeByRec(vals []int, idx int) *TreeNode {
	// 构建下标值为idx处的节点(如果vals[idx]==0 零值表示该节点不存在)
	if idx >= len(vals) || vals[idx] == 0 {
		return nil
	}

	// 初始化根节点
	root := &TreeNode{
		Val: vals[idx],
	}

	// 递归初始左右节点(2*i+1、2*i+2) -> 这个是完全二叉树的特性！
	root.Left = buildCompleteBinaryTreeByRec(vals, 2*idx+1)
	root.Right = buildCompleteBinaryTreeByRec(vals, 2*idx+2)

	return root
}

// 迭代方式构建二叉树
func buildBinaryTreeByLoop(vals []int) *TreeNode {
	treeSize := len(vals)
	if treeSize == 0 {
		return nil
	}

	idx := 1
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}

	var node *TreeNode
	for idx < treeSize {
		if len(queue) > 0 {
			node = queue[0]
			queue = queue[1:]
		}

		// 构建节点左节点
		if vals[idx] != 0 {
			node.Left = &TreeNode{Val: vals[idx]}
			queue = append(queue, node.Left)
		}
		idx++

		// 构建节点右子树
		if idx < treeSize && vals[idx] != 0 {
			node.Right = &TreeNode{Val: vals[idx]}
			queue = append(queue, node.Right)
		}
		idx++
	}

	return root
}
