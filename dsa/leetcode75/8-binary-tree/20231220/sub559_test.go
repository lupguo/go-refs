package _0231220

import (
	. "dsa/data-struct"
)

// 给定一个 N 叉树，找到其最大深度。
//
//	思路1: bfs
func maxDepthNTreeByBFS(root *Node) int {
	if root == nil {
		return 0
	}

	// bfs思想, que队列实现
	var level int
	que := []*Node{root}
	for len(que) > 0 {
		var lque []*Node
		for _, node := range que { // 遍历清空队列中所有节点
			lque = append(lque, node.Children...)
		}
		level++
		que = lque
	}

	return level
}

// 思路2: 先序遍历N叉树
func maxDepthNTreeByPreOrder(root *Node) int {
	var depth, maxNTreeDepth int
	return preOrderTraverseNTree(root, depth, maxNTreeDepth)
}

// 递归先序遍历N叉树
func preOrderTraverseNTree(root *Node, depth int, maxDepth int) int {
	if root == nil {
		return maxDepth
	}

	// 当前节点深度&最大深度
	depth++
	maxDepth = max(depth, maxDepth)
	for _, child := range root.Children { // 依次遍历子节点，看是否有超过最大深度的
		maxDepth = max(preOrderTraverseNTree(child, depth, maxDepth), maxDepth)
	}
	depth--

	return maxDepth
}
