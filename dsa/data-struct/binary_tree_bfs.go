package data_struct

// BFS 二叉树BFS 广度优先遍历（层次遍历）
// 思路: 将每层节点依次放入一个Queue队列中，通过range迭代直至为空
func (root *TreeNode) BFS() []int {
	if root == nil {
		return nil
	}

	// 非空
	var nodeVals []int
	nodeQueue := []*TreeNode{root}

	// 从根节点开始BFS迭迭代
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		nodeVals = append(nodeVals, node.Val)

		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
		}
	}

	return nodeVals
}
