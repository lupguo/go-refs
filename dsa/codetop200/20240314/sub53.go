package _0240314

import (
	"math"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6
//
// 思路，应该是要用到DP
// dp[i] = max(nums[i], nums[i]+dp[i-1])
// dp[i] 表示截止idx=i(包含)位置，的最大子数组和
// 问题关键几个点:
//  1. 想清楚如何得到递推公式（一定要想清楚dp动态规划函数的定义，以及状态转移方程）
//  2. 思考如何利用mem加速（在使用mem加速的时候，一定要清楚边界问题，举例、举例、举例）
func maxSubArray(nums []int) int {

	// mem[i]表示，包含nums[0..i](包含i下标)元素的最大子序列和（所以mem长度应该为n-1)
	mem := make([]int, len(nums))
	for i, _ := range mem {
		mem[i] = math.MinInt
	}

	// dp表示nums数组到n-1位置的最大子序列和，并借助mem加速
	return dp(nums, len(nums)-1, mem)
}

// 连续子数组，以end位置结束(包含)的最大子数组和
func dp(nums []int, idx int, mem []int) int {
	if nums == nil {
		return 0
	} else if idx == 0 {
		return nums[0]
	}

	// 缓存
	if mem[idx] != math.MinInt {
		return mem[idx]
	}

	// 递归求解
	res := math.MinInt
	for i := 0; i <= idx; i++ {
		if i == 0 {
			mem[i] = nums[i]
		} else {
			mem[i] = max(nums[i], nums[i]+dp(nums, i-1, mem))
		}
		res = max(mem[i], res)
	}

	return res
}
