package _0240223

import (
	"testing"
)

// https://leetcode.cn/problems/house-robber/description/
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 思路: 动态规划
//  1. 尝试从nums 任意选择一家，i in choices
//  2. 统计偷取金额，获取剩余选择 choices （stolen 排除掉邻居）
//  3. 递归rob剩余部分，直到没有可以偷取选择，统计这种方式偷取掉金额，和最值比较
func rob(nums []int) int {
	// 本轮偷取以及最大偷取金额
	var stolenMoney, maxStolen int
	stolen := make(map[int]bool)

	var dp func(choices []int, stolenMoney int, stolen map[int]bool)

	dp = func(nums []int, stolenMoney int, stolen map[int]bool) {
		// 无房可偷取，返还已经偷取掉总额
		if len(stolen) == len(nums) {
			if stolenMoney > maxStolen {
				maxStolen = stolenMoney
			}
			return
		}

		// 选择一种偷取策略
		for room, money := range nums {
			// 检测room能否偷取
			if stolen[room] {
				continue
			}

			// 计算该偷取策略偷取总额
			if room > 1 {
				stolen[room-1] = true
			}
			stolen[room] = true
			if room < len(nums)-1 {
				stolen[room+1] = true
			}
			stolenMoney += money

			// 刨除剩余掉
			dp(nums, stolenMoney, stolen)
			stolenMoney -= money

			if room > 1 {
				stolen[room-1] = false
			}
			stolen[room] = false
			if room < len(nums)-1 {
				stolen[room+1] = false
			}
		}
	}
	dp(nums, stolenMoney, stolen)

	return maxStolen
}

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1, 2, 3, 1}, 4},
		{"t2", []int{2, 7, 9, 3, 1}, 12},
		{"t3", []int{2, 1, 1, 2}, 4},
		{"t4", []int{5, 3, 4, 11, 2}, 16},
		{"t5", []int{1, 2, 3}, 4},
		{"t6", []int{1, 2, 1, 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
