package _0231204

import (
	"testing"

	. "dsa/data-struct"
)

// 统计树中好节点数目（从根到节点，均小于当前节点）
// 思路: 递归写法
//  1. 创建一个新的foundGoodNodes(root, max) []*TreeNode 查询root树拥有好节点数目
//  2. 申请一个goodNodes []*TreeNode记录好节点数量
//  3. 基准条件,
//     * root == nil, return0
//     * root !=nil , if root.Val >= max {
//     appnd到好节点中
//     更新max=root.Val
//     }
//  4. 如果root左子树非空，查询左子树 foundGoodNodes(root.Left, max)，同理右子树处理
//  5. 返回append(leftGoodNodes, rightGoodNodes)结果
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 比较节点和
	maxVal := root.Val
	goodNodes := []*TreeNode{root}

	// 递归求goodNodes数据
	goodNodes = findGoodNodes(root, maxVal)

	return len(goodNodes)
}

// 递归查询节点中好节点数量
func findGoodNodes(root *TreeNode, maxVal int) []*TreeNode {
	var nodes []*TreeNode

	// 查看当前节点是否满足好节点需求
	if root.Val >= maxVal {
		nodes = append(nodes, root)
		maxVal = root.Val
	}

	// 左节点好节点数据
	if root.Left != nil {
		nodes = append(nodes, findGoodNodes(root.Left, maxVal)...)
	}

	// 右节点中的好节点数据
	if root.Right != nil {
		nodes = append(nodes, findGoodNodes(root.Right, maxVal)...)
	}

	return nodes
}

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{"t1", []int{1}, 1},
		{"t2", []int{1, 2}, 2},
		{"t3", []int{2, 1}, 1},
		{"t4", []int{3, 1, 2}, 1},
		{"t5", []int{3, 1, 0, 2}, 1},
		{"t6", []int{1, 2, 3}, 3},
		{"t7", []int{1, 2, 3, 3}, 4},
		{"t8", []int{1, 1, 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootTree := IntSliceBFSToBinaryTree(tt.root)
			if got := goodNodes(rootTree); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
