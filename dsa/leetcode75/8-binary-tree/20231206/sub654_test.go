package _0231206

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
//
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
// 返回 nums 构建的 最大二叉树 。

// 思路: 递归思路，因为从列表中找到最大值，拆分成左右两个列表后，子问题也是同样问题
//  1. 查找nums的最大值 val，以及在列表中的索引位置maxIdx, 将数组分拆成两个子区间
//     - leftNums := nums[0:maxIdx],
//     - rightNums := nums[maxIdx+1:]
//  2. 构建节点root := &TreeNode{nums[maxIdx}
//     - root.Left = constructMaximumBinaryTree(leftNums)
//     - root.Right = constructMaximumBinaryTree(rightNums)
//  3. return root
func constructMaximumBinaryTree(nums []int) *TreeNode {
	lenNums := len(nums)
	if lenNums == 0 {
		return nil
	}

	// 查找最大值
	var maxIdx int
	for i := 0; i < lenNums; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	// 生成左、右nums区间
	leftNums, rightNums := nums[0:maxIdx], nums[maxIdx+1:]

	// 构建root节点 && 左右子树
	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTree(leftNums)
	root.Right = constructMaximumBinaryTree(rightNums)

	return root
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{3, 2, 1, 6, 0, 5}, []int{6, 3, 5, 2, 0, 1}},
		{"t2", []int{3, 2, 1}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTree := constructMaximumBinaryTree(tt.nums)
			gotBfsVals := gotTree.BFS()
			if !slices.Equal(gotBfsVals, tt.want) {
				t.Errorf("got %v, but want %v", gotBfsVals, tt.want)
			}
		})
	}
}
