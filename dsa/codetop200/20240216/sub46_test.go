package _0240216

import (
	"testing"
)

// 回溯算法，全排列问题
// https://leetcode.cn/problems/permutations/
//
//	思路: backtrack检测, 可行选择、做选择、递归backtrack检测
func permute(nums []int) [][]int {
	var res [][]int
	choices := nums
	var path []int

	// 递归回溯所有可行解
	backtrack(path, len(nums), choices, &res)
	return res
}

func backtrack(path []int, length int, choices []int, res *[][]int) {
	if len(path) == length {
		*res = append(*res, path) // 找到一个答案解
		return
	}

	// 选择列表
	for _, choice := range choices {
		// 做选择
		path = append(path, choice)
		reminder := removeElem(choices, choice)

		// 递归下层决策
		backtrack(path, length, reminder, res)

		// 选择回退
		path = removeElem(path, choice)
		choices = append(reminder, choice)
	}
}

// 删除元素
func removeElem(elems []int, remove int) []int {
	var ret []int
	for _, elem := range elems {
		if remove != elem {
			ret = append(ret, elem)
		}
	}
	return ret
}

func TestPermute(t *testing.T) {
	t.Logf("ret %v", permute([]int{1, 2, 3}))
}
