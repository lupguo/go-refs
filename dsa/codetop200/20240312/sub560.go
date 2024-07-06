package _0240312

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
//
// 方法1: 暴力算法，两层循环, sum(nums[i:j])==k count++
func subarraySumV1(nums []int, k int) int {
	var count int
	for i, _ := range nums {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// 方法2: 尝试用差值计算 sum(num[i:j]) = sum(0:j) - sum(0:i)
func subarraySum(nums []int, k int) int {

	// 初始化sum和
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] += sum[i-1] + nums[i-1]
	}

	// 区间差值查找k
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sum[j]-sum[i] == k {
				count++
			}
		}
	}

	return count
}
