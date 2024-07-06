package _0231214

import (
	"testing"

	. "dsa/data-struct"
)

var res int

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func longestZigZag(root *TreeNode) int {
	getPathLen(root)
	return res
}

// 计算并返回一个节点的左右交错路径长度，然后在后序位置上更新全局最大值。
func getPathLen(root *TreeNode) []int {
	if root == nil {
		return []int{-1, -1}
	}
	left := getPathLen(root.Left)
	right := getPathLen(root.Right)

	// 后序位置，根据左右子树的交错路径长度推算根节点的交错路径长度
	rootPathLen1 := left[1] + 1
	rootPathLen2 := right[0] + 1
	// 更新全局最大值
	res = max(res, max(rootPathLen1, rootPathLen2))

	return []int{rootPathLen1, rootPathLen2}
}

// 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
//
// 思路: 分治+递归，分别求左右子树的最大交错路径，取最大值，得到结果
//  1. 定义一个新的函数，记录从zigZag指定方向得到到的最大值，zigZagRange(node, isLeft) int
//  2. longestZigZagRec(root) = max(zigZagRange(node.left, true), zigZagRange(node.right, false))
func longestZigZagRec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 当前节点交叉深度
	maxPath := max(zigZagRange(root.Left, true), zigZagRange(root.Right, false))

	// 左、右节点交叉深度
	leftMaxPath := longestZigZag(root.Left)
	rightMaxPath := longestZigZag(root.Right)

	// 返回左、右、当前节点深度
	return max(maxPath, leftMaxPath, rightMaxPath)
}

// 记录从指定节点开始，取得的最大路值
func zigZagRange(node *TreeNode, fromLeft bool) int {
	if node == nil {
		return 0
	}
	// 若从左边过来的，下次遍历到右侧，反之一样
	if fromLeft == true {
		return 1 + zigZagRange(node.Right, false)
	} else {
		return 1 + zigZagRange(node.Left, true)
	}
}

func TestLongestZigZag(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  int
	}{
		{"t1", []int{1}, 0},
		{"t2", []int{1, 2}, 1},
		{"t3", []int{1, 2, 3}, 1},
		{"t4", []int{1, 2, 3, 4}, 1},
		{"t5", []int{1, 2, 3, 4, 5}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tTree := IntSliceBFSToBinaryTree(tt.nodes)
			if got := longestZigZag(tTree); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
