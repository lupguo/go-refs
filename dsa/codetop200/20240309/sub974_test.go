package _0240309

import (
	"testing"
)

func subarraysDivByK(nums []int, k int) int {
	// base case
	if k == 0 {
		return 0
	}

	// 计算前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 重新迭代数组，依次检测从 nums[i] -> nums[j]期间有多少可能到解
	count := 0
	for i := 0; i < len(nums); i++ {
		// 内层循环检测，nums[i:j]子数组和是否满足k的整除
		for j := i + 1; j <= len(nums); j++ {
			if (preSum[j]-preSum[i])%k == 0 {
				count++
			}
		}
	}

	return count
}

func TestSubarraysDivByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"t1", []int{1}, 1, 1},
		{"t2", []int{1}, 0, 0},
		{"t3", []int{1, 2, 3}, 3, 3},
		{"t4", []int{1, 2, 3, 2}, 5, 2},
		{"t5", []int{1, 2, 3, 2, -7}, 5, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ans := subarraysDivByKV2(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("got %v, but want %v, got ans %v", got, tt.want, ans)
			}
		})
	}
}

// https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/
// 给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
//
// 子数组 是数组的 连续 部分。
// 思路:
//  1. 因为问题要求是连续的子数组，且涉及到连续子数组到和，[a] [a,b], [a,b,c]..，想到前缀和方式求解
//  2. 先申请数组长度大小n+1，存储每个到下标位置到元素和preSum
//  3. 依次迭代数组i，内层j循环检测，nums[i:j]子数组和是否满足k的整除
func subarraysDivByKV2(nums []int, k int) (int, [][]int) {
	if k == 0 {
		return 0, nil
	}
	// 计算前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 重新迭代数组，依次检测从 nums[i] -> nums[j]期间有多少可能到解
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if (preSum[j]-preSum[i])%k == 0 {
				ans = append(ans, nums[i:j])
			}
		}
	}

	return len(ans), ans
}
