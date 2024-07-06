package _0231220

import (
	"testing"

	. "dsa/data-struct"
)

// 给定一个二叉树，找出其最小深度
//
// 思路：getMinPath(root) = min(travser(root.left), travser(root.right))
//   - 遍历过程，借助全局的depth 和 minDepth进行更新
func minDepth(root *TreeNode) int {
	traverseMinDepth(root)
	return minDepthPath
}

var depth, minDepthPath int

func traverseMinDepth(root *TreeNode) {
	if root == nil {
		return
	}

	// 先序遍历访问位置
	depth++

	// 叶子节点视情况更新最新深度
	if root.Left == nil && root.Right == nil {
		if minDepthPath > 0 {
			minDepthPath = min(depth, minDepthPath)
		} else {
			minDepthPath = depth
		}
	}

	traverseMinDepth(root.Left)
	traverseMinDepth(root.Right)

	// 后续遍历离开节点
	depth--
}

// 给定一个二叉树，找出其最小深度
// 思路: 树的最小深度 getMinPath(root) = min(getMinPath(root.left), getMinPath(root.right))
func minDepthV3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth, minPath := 0, 0
	return getMinPath(root, depth, minPath)
}

func getMinPath(root *TreeNode, depth, minPath int) int {
	if root == nil {
		return minPath
	}

	// 节点深度+1
	depth++
	// 当前节点是叶子节点，更新最小深度，或者返回已知叶子节点最小深度
	if root.Left == nil && root.Right == nil {
		if minPath == 0 {
			minPath = depth
		}
		return min(depth, minPath)
	}

	// 非叶子节点，且当前节点深度已比最小节点深度还大，直接剪枝返回
	if minPath > 0 && depth >= minPath {
		return minPath
	}

	// 非叶子节点，递归找左子树最小深度、右子树找最小深度
	return getMinPath(root.Right, depth, getMinPath(root.Left, depth, minPath))
}

// 给定一个二叉树，找出其最小深度
// 思路: 遍历位置，统计到叶子节点(left,right==nil)的path长度，超过了最小值则直接剪枝
func minDepthV2(root *TreeNode) int {
	var depth, minTreeDepth int
	return traverseTree(root, depth, minTreeDepth)
}

func traverseTree(root *TreeNode, depth int, minDepthPath int) int {
	if root == nil {
		return minDepthPath
	}

	// 先序遍历位置，进入节点，深度+1 && 剪枝
	depth++
	if minDepthPath > 0 && depth > minDepthPath {
		depth--
		return minDepthPath
	}

	// 叶子节点更新minDepth
	if root.Left == nil && root.Right == nil {
		if minDepthPath > 0 {
			minDepthPath = min(depth, minDepthPath)
		} else {
			minDepthPath = depth
		}
		depth--
		return minDepthPath
	}

	// 递归子树获取最小深度
	minDepthPath = traverseTree(root.Left, depth, minDepthPath)
	minDepthPath = traverseTree(root.Right, depth, minDepthPath)

	// 回溯时候depth要-1
	depth--

	return minDepthPath
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  int
	}{
		{"t1", []int{1, 2, 3, 4}, 2},
		{"t2", []int{1, 2}, 2},
		{"t3", []int{1, 2, 0, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nodes)
			if got := minDepth(root); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
