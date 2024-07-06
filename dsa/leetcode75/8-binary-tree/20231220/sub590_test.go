package _0231220

import (
	. "dsa/data-struct"
)

// N叉树后续遍历
// 注意危险: 切片的全局危险使用
func postorder(root *Node) []int {
	var ret []int
	if root == nil {
		return nil
	}

	for _, child := range root.Children {
		ret = append(ret, postorder(child)...)
	}

	ret = append(ret, root.Val)
	return ret
}

// N叉树先续遍历
//
//	思路1: 根 -> range{ preorder(root.child)}
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int

	ret = append(ret, root.Val)
	for _, child := range root.Children {
		ret = append(ret, preorder(child)...)
	}

	return ret
}

// 思路2：BFS方式实现先序
func preorderBFS(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	que := []*Node{root}

	// bfs队列非空
	for len(que) > 0 {
		// 出队&
		node := que[0]
		que = que[1:]
		ret = append(ret, node.Val)

		var lque []*Node
		for _, node := range que {
			lque = append(lque, node.Children...)
		}
		que = lque
	}

	return ret
}
