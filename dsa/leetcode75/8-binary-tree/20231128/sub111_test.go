package _0231128

import (
	"testing"

	. "dsa/data-struct"
)

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
// 思路: 通过递归方式求解，分清楚数节点情况
//  1. root为空，返回0
//  2. root为子节点，返回1
//  3. 左、右子树存在一个空节点:
//     * root左节点为空，返回 minDepth(root.Right)+1
//     * root右节点为空，返回 minDepth(root.Left)+1
//  4. 左右节点均非空: 返回 min(minDepth(root.Right), minDepth(root.Left)) +1
func minDepth2(root *TreeNode) int {
	// 为子节点，返回1
	switch {
	case root == nil:
		return 0
	case root.Left == nil && root.Right == nil:
		return 1
	case root.Left == nil:
		return minDepth(root.Right) + 1
	case root.Right == nil:
		return minDepth(root.Left) + 1
	default:
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

// 思路2: 采用非递归，通过BFS方式，一层层统计，每层深度+1,最先在某一层找到子节点，则返回该depth
//  1. 判断root是否非空，空返回0
//  2. 申请depth, 和一个[]*TreeNode队列（队列中的元素为当前层的所有元素，需要一次遍历完成），初始root第一次队列，用于BFS
//  3. 判断队列长度，如果队列长度 > 0 ，则depth++，需要一次遍历队列中所有元素
//  4. range queue,
//     * 判断是否存在子节点，有则返回depth为最小深度
//     * 否则将得到的节点存入levelNodes
//  5. range当前层结束后，将levelNodes元素加入到queue中，BFS下一层
func minDepth(root *TreeNode) int {
	// 异常情况
	if root == nil {
		return 0
	}

	// 层次遍历
	queueNodes := []*TreeNode{root}
	var depth int

	// 迭代层queue
	for len(queueNodes) > 0 {
		depth++
		var levelNodes []*TreeNode
		for _, node := range queueNodes {
			// node是否为子节点
			if node.Left == nil && node.Right == nil {
				return depth
			} else if node.Right == nil {
				levelNodes = append(levelNodes, node.Left)
			} else if node.Left == nil {
				levelNodes = append(levelNodes, node.Right)
			} else {
				levelNodes = append(levelNodes, node.Left, node.Right)
			}
		}

		// 层次遍历完成，重新入队
		queueNodes = levelNodes
	}

	return depth
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{"t1", []int{1}, 1},
		{"t2", []int{1, 2}, 2},
		{"t3", []int{1, 2, 3, 4}, 2},
		{"t4", []int{1, 2, 3, 0, 4, 5, 6, 0, 0, 0, 7}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTree := IntSliceBFSToBinaryTree(tt.root)
			if got := minDepth2(testTree); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
