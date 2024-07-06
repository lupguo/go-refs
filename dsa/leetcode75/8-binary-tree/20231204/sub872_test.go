package _0231204

import (
	"testing"

	. "dsa/data-struct"
)

// 判断两颗二叉树，是否叶子序相同(叶相似）如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的
//
// 思路1：基于中序遍历，得到两颗树的叶子节点数组，依次判断两个叶子节点序数组是否一致
//  1. 申请两个叶子节点序切片
//  2. 中序遍历两棵树，得到叶子节点的序内容
//  3. 比较两颗树的叶子节点序是否一致
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leafs1, leafs2 []*TreeNode

	// 中序遍历，得到叶子序内容
	leafs1 = getLeafNodes(root1)
	leafs2 = getLeafNodes(root2)

	// 比较两颗叶子序是否一致
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i, _ := range leafs1 {
		if leafs1[i].Val != leafs2[i].Val {
			return false
		}
	}

	return true
}

// 基于中序遍历，获取一颗树的叶子节点序
func getLeafNodes(root *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	// 子节点返回
	if root.Left == nil && root.Right == nil {
		return []*TreeNode{root}
	}

	// 递归获取左右子树的子节点信息
	if root.Left != nil {
		nodes = append(nodes, getLeafNodes(root.Left)...)
	}

	if root.Right != nil {
		nodes = append(nodes, getLeafNodes(root.Right)...)
	}

	return nodes
}

// 获取相似子树
func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		name  string
		root1 []int
		root2 []int
		want  bool
	}{
		{"t1", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"t2", []int{1, 2, 3}, []int{1, 3, 2}, false},
		{"t3", []int{1, 0, 2}, []int{1, 2}, true},
		{"t4", []int{1, 0, 2, 0, 0, 3, 4}, []int{1, 2, 4, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r1 := IntSliceBFSToBinaryTree(tt.root1)
			r2 := IntSliceBFSToBinaryTree(tt.root2)
			if got := leafSimilar(r1, r2); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
