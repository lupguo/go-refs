package _0240223

import (
	"testing"

	. "dsa/data-struct"
)

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
//
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//		思路: 因为是BST树，从小到大第k个 = 中序遍历+计数方式解决
//	 基于BST性质，我们可以使用中序遍历（Inorder Traversal）来得到一个升序排列的数组。因此，要找到第 k 个最小元素，只需要对 BST 进行中序遍历，并记录访问过程中经过的元素数量即可。
func kthSmallest(root *TreeNode, k int) int {
	count := 0 // 记录已经访问过多少个元素
	found := false
	var result int

	// 定义inorder 闭包
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		// base
		if node == nil {
			return
		}
		// 递归左子树
		inorder(node.Left)

		// 中序位置
		count++
		if count == k { // 找到第 k 小元素了
			result = node.Val
			found = true
			return
		}
		if found {
			return
		}

		// 递归右子树
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func TestKthSmallest(t *testing.T) {
	root := IntSliceBFSToBinaryTree([]int{3, 1, 4, 0, 2})
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{"t1", root, 1, 1},
		{"t2", root, 2, 2},
		{"t3", root, 3, 3},
		{"t4", root, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
