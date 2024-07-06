package _0231214

import (
	"testing"

	. "dsa/data-struct"
)

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//	思路: 最近公共祖先
//	    1. 如果p,q分别处于root节点的左、右子树中，那么root则肯定为p、q的最近公共祖先
//	    2. 因为p,q给到均在树中，所以如果p,q不在root节点左、右子树上，则一定在子树的同侧，最先找到的节点即为最近公共祖先节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 从子树找到了一个值
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 左、右子树找节点数
	leftFoundNode := lowestCommonAncestor(root.Left, p, q)
	rightFoundNode := lowestCommonAncestor(root.Right, p, q)

	// 从root的左、右子树都找到了一个值，那么表示root就是p,q的最近公共祖先
	if leftFoundNode != nil && rightFoundNode != nil {
		return root
	} else if leftFoundNode != nil { // 仅一边找到
		return leftFoundNode
	} else {
		return rightFoundNode
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		p, q  int
		want  int
	}{
		{"t1", []int{1, 2, 3, 4, 5}, 1, 5, 1},
		{"t2", []int{1, 2, 3, 4, 5}, 2, 3, 1},
		{"t3", []int{1, 2, 3, 4, 5}, 3, 4, 1},
		{"t4", []int{1, 2, 3, 4, 5}, 2, 4, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := IntSliceBFSToBinaryTree(tt.nodes)
			p, q := &TreeNode{Val: tt.p}, &TreeNode{Val: tt.q}
			got := lowestCommonAncestor(root, p, q)
			if got == nil || got.Val != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
