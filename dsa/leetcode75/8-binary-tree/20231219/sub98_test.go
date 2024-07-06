package _0231219

import (
	"math"
	"testing"

	. "dsa/data-struct"
)

var maxNum int

// 验证一棵树是否是BST二叉搜索树
// 思路:
//  1. 通过中序遍历后，BST满足之前遍历的最大值一定都会小于当前值
func isValidBST(root *TreeNode) bool {
	maxNum = math.MinInt64
	if ok := matchInorderSort(root); ok {
		return true
	}
	return false
}

func matchInorderSort(root *TreeNode) (isMatch bool) {
	if root == nil {
		return true
	}

	// 迭代左子树
	ok := matchInorderSort(root.Left)
	if !ok {
		return false
	}

	// BST根节点应该比之前迭代的maxNum大，否则不满足BTS二叉树特性
	curNum := root.Val
	if maxNum >= curNum {
		return false
	}
	maxNum = curNum

	// 迭代右子树也满足二叉树特性
	ok = matchInorderSort(root.Right)
	if !ok {
		return false
	}

	return true
}

// 思路2：
//   - 递归的检测是否满足BST特性
//   - 节点值要大于节点的左子树所有元素（root作为maxNode传入）
//   - 节点值要小于节点的右子树所有元素（root作为minNode传入）
//     通过min.val < root.val < max.val检测
func isValidBSTV2(root *TreeNode) bool {
	return matchBST(root, nil, nil)
}

// minNode表示当子树的最小值，
func matchBST(root *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if root == nil {
		return true
	}

	// 按理 minNode <= root <= maxNode
	if minNode != nil && root.Val <= minNode.Val {
		return false
	}
	if maxNode != nil && root.Val >= maxNode.Val {
		return false
	}

	// 递归检测左右子树是否符合二叉
	return matchBST(root.Left, nil, root) && matchBST(root.Right, root, nil)
}

// minNode表示当子树的最小值，
func matchBST1(root *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if root == nil {
		return true
	}

	// 检测是否满足BST特性，左子树根节点值比最小值还小，不满足
	if minNode != nil && root.Val <= minNode.Val {
		return false
	}
	if maxNode != nil && root.Val >= maxNode.Val {
		return false
	}

	// 递归检测root的左右子树是否满足特性
	return matchBST(root.Left, minNode, root) && matchBST(root.Right, root, maxNode)
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  bool
	}{
		{"t1", []int{2, 1, 3}, true},
		{"t2", []int{2, 1}, true},
		{"t3", []int{2, 0, 3}, true},
		{"t4", []int{1, 2, 3}, false},
		{"t5", []int{1, 3, 2}, false},
		{"t6", []int{3, 2, 4, 1, 0, 5}, false},
		{"t7", []int{3, 2, 4, 1, 0, 0, 5}, true},
		{"t8", []int{2, 2, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nodes)
			if got := isValidBSTV2(root); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
