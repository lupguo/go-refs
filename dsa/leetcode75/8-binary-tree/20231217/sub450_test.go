package _0231217

import (
	"testing"

	. "dsa/data-struct"
)

// 从二叉搜索树BST中删除掉某个节点，返回删除后新的BST
//
//	思路:
//	    1. 先序遍历，找到要删除的节点node
//	    2. 基于BST性质：
//	        * node key的右子树非空, node左子节点值 < node值 < node右子节点值，需要找到node节点右子树中最左叶子节点
//	        * node key的右子树为空，将node的左子树，接入到node key 父节点的左节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// root 根节点是否要删除的元素，
	if root.Val == key {
		// 删除元素的右子树非空，从右子树找到最小节点
		if root.Right != nil {
			rTreeMinNode := removeBSTMinNode(root.Right)
			rTreeMinNode.Left = root.Left
			rTreeMinNode.Right = root.Right
			return rTreeMinNode
		} else {
			return root.Left
		}
	}

	// 如果root不是要删除的元素，递归的从左、右子树删除
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

// 移除BST最小元素(一定是叶子节点)并删除
//
//	思路:
//	    1. root就一个节点，返回空
func removeBSTMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	p := root.Left
	if p == nil {
		return root
	}

	// 左子树节点就是要删除的元素
	var delNode *TreeNode
	for root.Left != nil {
		// root.left就是要删除的元素
		if root.Left.Val == key {
			delNode = root.Left
			root.Left = nil
		}
		root = root.Left
	}

	return delNode
}

func TestRemoveBSTMinNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		key   int
		want  []int
	}{
		{"t1", []int{2, 1, 3}},
	}
}

func TestDeleteNode(t *testing.T) {

}
