package dfs

import (
	. "dsa/data-struct"
)

// 给定一个二叉树 root ，返回其最大深度。
func maxDepthStack(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}

	// 递归找左右子树高度
	leftNodeHeight := maxDepth(root.Left)
	rightNodeHeight := maxDepth(root.Right)

	// 找到一个最大的
	if leftNodeHeight > rightNodeHeight {
		return leftNodeHeight + 1
	}

	return rightNodeHeight + 1
}

// 给定一个二叉树 root ，返回其最大深度。
// 思路: 递归思路, max(maxDepth(root.Left), maxDepth(root.Right))+1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// 思路: 非递归，通过迭代方式编写- BFS思路
// 1. 从根节点开始，初始化装入queue队列
// 2. 不断遍历queue中的内容，直至为空
// 3. 每次遍历完整遍历迭代每一层的节点，迭代完后层数加1
// 4. 将每层的节点的非空左、右子节点，加入到queue中
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 层数
	depth := 1
	queue := []*TreeNode{root}
	for {
		// 完整将一层遍历完后，将下层非空节点重新放入队列技术，深度加一
		var subNodes []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				subNodes = append(subNodes, node.Left)
			}
			if node.Right != nil {
				subNodes = append(subNodes, node.Right)
			}
		}

		// 如果子节点加入到队列后，队列非空，则层深度+1，重新迭代子队列内容
		if len(subNodes) != 0 {
			depth++
			queue = subNodes
			continue
		}

		return depth
	}

}
