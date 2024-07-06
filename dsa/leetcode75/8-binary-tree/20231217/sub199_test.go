package _0231217

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// 从右侧看二叉树
//
//	思路: 利用BFS层次遍历二叉树方式，每次将层的最后一个节点值取出存入后，最后统一返回
//	    1. BFS一颗二叉树，使用Que队列
//	    2. 层次遍历需要一次将Que内迭代完
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// BFS
	queue := []*TreeNode{root}
	var ret []int
	for len(queue) > 0 {
		var children []*TreeNode
		for i, node := range queue {
			// 将node的子节点统一加入到childs中
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}

			// 迭代队列尾部元素
			if i == len(queue)-1 {
				ret = append(ret, node.Val)
			}
		}

		// 继续下一层
		queue = children
	}

	return ret
}

// 思路2: 利用DFS思路(根->右子节点->左子节点)，每层最先被访问到的节点就是ret结果集的元素
func rightSideViewByDFS(root *TreeNode) []int {
	var ret []int
	var depth int
	return dfs(root, depth, ret)
}

// dfs遍历一颗树
func dfs(root *TreeNode, depth int, ret []int) []int {
	if root == nil {
		return ret
	}

	// 根节点处理，如果在depth层元素未添加，则添加
	if depth == len(ret) {
		ret = append(ret, root.Val)
	}

	// dfs右子树，再dft左子树
	depth++
	ret = dfs(root.Right, depth, ret)
	ret = dfs(root.Left, depth, ret)

	return ret
}

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  []int
	}{
		{"t1", []int{1, 2, 3}, []int{1, 3}},
		{"t2", []int{1, 2, 3, 4}, []int{1, 3, 4}},
		{"t3", []int{1, 0, 3, 0, 0, 4}, []int{1, 3, 4}},
		{"t4", []int{1, 0, 3}, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nodes)
			got := rightSideViewByDFS(root)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
