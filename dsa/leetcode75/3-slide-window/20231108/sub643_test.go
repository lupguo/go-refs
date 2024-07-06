package _0231108

import (
	"math"
	"testing"
)

// 题解: 找连续k个子数组，要求其该子数组平均数最大

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
// 任何误差小于 10^-5 的答案都将被视为正确答案。
//
// 示例 1：
//
// 输入：nums = [1,12,-5,-6,50,3], k = 4
// 输出：12.75
// 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
// 示例 2：
//
// 输入：nums = [5], k = 1
// 输出：5.00000
func findMaxAverage(nums []int, k int) float64 {

	return 0
}

func Method4FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)

	// 统计前k个数的sum
	var curSum int
	for i := 0; i < k; i++ {
		curSum += nums[i]
	}
	maxSum := curSum

	// 从k位开始，滑动统计每k个数的sum
	// curSum表示从[i-k+1, i]位置的sum
	// curSum[i-k+1] = curSum[i-k] - nums[i-k] + num[i]
	for i := k; i < n; i++ {
		curSum = curSum - nums[i-k] + nums[i]
		maxSum = max(curSum, maxSum)
	}

	return float64(maxSum) / float64(k)
}

// 方法2: 滑动窗口移动优化
// i从0到n, 窗口移动一格，新增新的数据，扣减老的数据，得到新的连续子序列和，将该子序列和和最大子序列和比较，返回最大子序列和
func Method3FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)

	var curSum int
	maxCnt := math.MinInt64
	for i := 0; i < n; i++ {
		curSum += nums[i]

		// 从k位开始
		if i >= k {
			curSum -= nums[i-k]
		}

		// 要从k-1才开始统计curSum,MaxCnt P.S，该代码实现方式，可读性还是比较差
		if i >= k-1 {
			maxCnt = max(curSum, maxCnt)
		}
	}

	return float64(maxCnt) / float64(k)
}

// 方法2: 连续K个子数组最大平均（和最大) => 寻找和最大的子数组
// => 规律 [i, i+k-1] = sum(i+1) = sum(i)+nums[i+k-1]-num[i] => 滑动窗口
func Method2FindMaxAverage(nums []int, k int) float64 {
	if k == 0 {
		return 0
	}

	// 初始化第一个窗口的子序列数组和，maxSum
	n := len(nums)
	var maxSum int
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	// 从i=1开始，迭代到n-k+1
	lastSum := maxSum
	for i := 1; i < n-k+1; i++ {
		curSum := lastSum + nums[i+k-1] - nums[i-1]
		if curSum > maxSum {
			maxSum = curSum
		}
		lastSum = curSum
	}

	return float64(maxSum) / float64(k)
}

// 从nums中找到k个数的连续子数组，计算其最大的平均值
// i从(0,n-k+1)，计算sum(num[i]...num[i+k-1])/k的最大值
// 暴力算法
func Method1FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	var maxAvg float64
	maxAvg = math.Inf(-1)
	for i := 0; i < n-k+1; i++ {

		var curSum float64
		for j := i; j < i+k; j++ {
			curSum += float64(nums[j])
		}

		if curAvg := curSum / float64(k); curAvg > maxAvg {
			maxAvg = curAvg
		}
	}

	return maxAvg
}

func TestMethod1FindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"t1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"t2", []int{5}, 1, 5},
		{"t3", []int{-1}, 1, -1.0},
		{"t4", []int{1}, 1, 1.0},
		{"t5", []int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{"t6", []int{1, 2, 3}, 3, 2},
		{"t7", []int{1, 2, 3, 4, 5}, 3, 4},
		{"t8", []int{-1, 2, -3, 4, -5}, 3, 1},
		{"t9", []int{-6662, 5432, -8558, -8935, 8731, -3083, 4115, 9931, -4006, -3284, -3024, 1714, -2825, -2374, -2750, -959, 6516, 9356, 8040, -2169, -9490, -3068, 6299, 7823, -9767, 5751, -7897, 6680, -1293, -3486, -6785, 6337, -9158, -4183, 6240, -2846, -2588, -5458, -9576, -1501, -908, -5477, 7596, -8863, -4088, 7922, 8231, -4928, 7636, -3994, -243, -1327, 8425, -3468, -4218, -364, 4257, 5690, 1035, 6217, 8880, 4127, -6299, -1831, 2854, -4498, -6983, -677, 2216, -1938, 3348, 4099, 3591, 9076, 942, 4571, -4200, 7271, -6920, -1886, 662, 7844, 3658, -6562, -2106, -296, -3280, 8909, -8352, -9413, 3513, 1352, -8825},
			90, 37.25556},
	}

	// t.Logf("1e-6=%v", 1e-6)
	// t.Logf("1e-5=%v", 1e-5)
	// t.Logf("1e-4=%v", 1e-4)
	// t.Logf("1e-3=%v", 1e-3)
	// t.Logf("1e-2=%v", 1e-2)
	// t.Logf("1e-1=%v", 1e-1)
	// t.Logf("1e-1=%v, %08[1]b", 1e-1)
	// t.Logf("1e0=%v, %[1]f", 1e0)
	// t.Logf("1e1=%v, %[1]f", 1e1)
	// t.Logf("1e2=%v, %[1]f", 1e2)
	// t.Logf("1e3=%v, %[1]f", 1e3)
	// t.Logf("1e3=%v, %[1]f", 1e3)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method3FindMaxAverage(tt.nums, tt.k); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("got %v, but want %v, math.Abs(got-tt.want)=%v", got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}
