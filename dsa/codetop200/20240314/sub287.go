package _0240314

import (
	"sort"
)

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
//
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间
// 思路: O(1) 的额外空间，所以这里不能申请额外数组
// 可以先排序 => [2,1,2] => [1,2,2] i游标迭代-> [1,2,2] nums[i]==nums[i+1] return
func findDuplicateV1(nums []int) int {
	// 先排序
	sort.Ints(nums)

	// 游标迭代
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}

// 思路2: 通过交换方式
// [2,1,2] => nums[i]=2，放到index=2位置，如果index=2 num[index]==nums[i]，则表述有冲突
func findDuplicate(nums []int) int {
	// n+1
	idx := nums[0] // 游标

	// 先暂存nums[idx]
	for idx != nums[idx] {
		tmp := nums[idx]
		nums[idx] = idx
		idx = tmp
	}

	return idx
}
