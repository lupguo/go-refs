package _0240314

import (
	"math"

	. "dsa/data-struct"
)

// 递归，直接求和
func sumNumbers(root *TreeNode) int {
	// 当前节点和
	return traverseSum(root, 0)
}

// 求到root节点的sum和
func traverseSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	// 当前节点数字
	sum = sum*10 + root.Val // 1

	// 抵达根节点
	if root.Left == nil && root.Right == nil {
		return sum
	}

	// 递归求解左右路径和
	return traverseSum(root.Left, sum) + traverseSum(root.Right, sum)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点
// 思路：dfs方法
func sumNumbersV1(root *TreeNode) int {

	var path []int // 从root到子节点的path元素
	var nums []int // 从root到子节点path元素数字

	// dfs遍历数组
	dfs(root, &path, &nums)

	// 求和
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}

// dfs遍历数组
func dfs(root *TreeNode, path *[]int, nums *[]int) {
	// base case
	if root == nil {
		return
	}

	// 将当前节点值加入path
	*path = append(*path, root.Val)

	// 是否抵达叶子节点，加入nums
	if root.Left == nil && root.Right == nil {
		// path[1,2,3] => 123
		*nums = append(*nums, pathToNum(*path))

		// 左右递归完成后，path要移除最后添加元素返回
		*path = (*path)[:(len(*path) - 1)]

		return
	}

	// 左、右子树非空，递归处理
	if root.Left != nil {
		dfs(root.Left, path, nums)
	}
	if root.Right != nil {
		dfs(root.Right, path, nums)
	}

	// 左右递归完成后，path要移除最后添加元素返回
	*path = (*path)[:(len(*path) - 1)]

	return
}

// 切片数字转整型数字
func pathToNum(numbs []int) int {
	var ret int
	n := len(numbs)
	for i := 0; i < n; i++ {
		ret += numbs[i] * int(math.Pow10(n-1-i))
	}
	return ret
}
