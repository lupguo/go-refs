package week17

import (
	"slices"
)

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 (重复也算，比如0,0)
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 思路:
//  1. 排序ONLog(N)，统计自增最长连续串，比如[0,1,3,4,5] => 3
//  2. 直接把Nums放入 tmp（0,N)到的数组中(空间O(max(N))，数组值为idx计数个数，比如1为下标tmp[1]++，再遍历数组[0,1,2,3,4,....100...200]
//     - 能直接通过交换么？将tmp[idx]位置移动到目标位置
func longestConsecutive(nums []int) int {
	// 将nums放入tmp数组
	minNum, maxNum := slices.Min(nums), slices.Max(nums)
	var added int
	if minNum < 0 {
		added = -minNum
	}

	// 全0情况
	if maxNum == 0 && minNum == maxNum {
		return len(nums)
	}

	// 将nums放入tmp数组中
	newNums := make([]int, maxNum+added)
	for _, num := range nums {
		idx := num + added
		newNums[idx]++
	}

	// 统计newNums连续最长子数组
	var ans, cnt int
	for _, num := range newNums {
		if num > 0 { // 刷新ans
			cnt += num
			ans = max(cnt, ans)
		} else {
			cnt = 0
		}
	}

	return ans
}
