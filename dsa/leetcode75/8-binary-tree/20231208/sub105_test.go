package _0231208

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 思路: 画图分析前序、中序规律，发现
//   - 可以通过前序构建根节点, 到中序找到根节点，得到左右子树中序序列，再回到前序得到左右子树前序序列，从而递归可以得到根的左、右子树
//   - O(logN)从中序遍历找根，分成 [左..] - 根 - [右..] 模式，递归得到树
//     1. base case: len(preorder) == 0 , return nil
//     2. rootVal := preOrder[0]   // 前序构建根节点
//     3. 在inorder中找rootVal值，得到rootValIdx
//     4. 将inoder分成两个子树的inorder区间:
//          - leftSubTreeInoder, rightSubTreeInorder := inorder[0:rootValIdx],  inorder[rootValIdx+1:]
//     5. 获取左、右子树节点数，leftSubTreeInoderLen, rightSubTreeIndoerLen
//     6. 将preorder分成两个子树的preorder区间:
//          - if leftSubTreeInoderLen > 0 { rightSubtreePreoder = preorder[1:leftSubTreeInoderLen] }
//          - if rightSubTreeInorderLen > 0 { rightSubTreePreorder = preorder[1+leftSubTreeInoderLen:]
//     8. 构建root的左右子树，并返回结果
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 根节点值
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 根节点值在中序遍历的位置
	rootValIdx := slices.Index(inorder, rootVal)

	// 将中序遍历分拆成两棵子树的中序遍历结果
	leftSubtreeInorder, rightSubtreeInorder := inorder[0:rootValIdx], inorder[rootValIdx+1:]

	// 左右子树先序遍历结果
	var leftSubtreePreorder, rightSubtreePreorder []int
	
	leftSubtreeInorderLen := len(leftSubtreeInorder)
	rightSubtreeInorderLen := len(rightSubtreeInorder)

	// 左树中序数组
	if leftSubtreeInorderLen > 0 {
		leftSubtreePreorder = preorder[1 : leftSubtreeInorderLen+1]
	}

	// 右树中序数组
	if rightSubtreeInorderLen > 0 {
		rightSubtreePreorder = preorder[1+leftSubtreeInorderLen:]
	}

	// 借助buildTree递归构建左右子树
	root.Left = buildTree(leftSubtreePreorder, leftSubtreeInorder)
	root.Right = buildTree(rightSubtreePreorder, rightSubtreeInorder)

	return root
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name    string
		preoder []int
		inorder []int
		want    []int
	}{
		{"t1", []int{1, 2, 4, 5, 3, 6}, []int{4, 2, 5, 1, 3, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"t2", []int{1, 2}, []int{2, 1}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTree := buildTree(tt.preoder, tt.inorder)
			gotTreeBFS := gotTree.BFS()
			if !slices.Equal(gotTreeBFS, tt.want) {
				t.Errorf("got %v, but want %v", gotTreeBFS, tt.want)
			}
		})
	}
}
