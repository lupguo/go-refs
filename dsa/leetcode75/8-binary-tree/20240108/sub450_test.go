package _0240108

import (
	"slices"
	"testing"

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

// 思考框架
//
//  1. 找到要删除的节点
//  2. 删除它
//
// 删除BST树(有序树)中某个节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 基本条件
	if root == nil {
		return nil
	}

	// 找到指定节点，基于root所在位置删除该节点
	if key == root.Val {

		// a. root为叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// b. root只有一个的左节点或右节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// c. root的左、右节点都不空，从右子树找到最小值，再将该值替换掉原先的root左、右节点
		rightMinChild := findDelRightMinChild(root.Right)
		rightMinChild.Right = deleteNode(root.Right, rightMinChild.Val)
		rightMinChild.Left = root.Left
		return rightMinChild

	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func findDelRightMinChild(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	// var parent *TreeNode
	for root.Left != nil {
		// parent = root
		root = root.Left
	}

	// parent.Left = nil
	return root
}

func TestDeleteBSTNode(t *testing.T) {
	tests := []struct {
		name     string
		nodeList []int
		key      int
		want     []int
	}{
		{"t1", []int{7, 4, 8, 3, 5, 0, 9}, 3, []int{7, 4, 8, 5, 9}},
		{"t2", []int{7, 4, 8, 3, 5, 0, 9}, 5, []int{7, 4, 8, 3, 9}},
		{"t3", []int{7, 4, 8, 3, 5, 0, 9}, 8, []int{7, 4, 9, 3, 5}},
		{"t4", []int{7, 4, 8, 3, 5, 0, 9}, 4, []int{7, 5, 8, 3, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTree := IntSliceBFSToBinaryTree(tt.nodeList)
			gotTree := deleteNode(testTree, tt.key)
			got := gotTree.BFS()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
