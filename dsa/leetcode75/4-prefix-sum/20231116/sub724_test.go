package _0231116

import (
	"testing"
)

// 寻找数组的中心下标
// 给你一个整数数组 nums ，请计算数组的 中心下标
//
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
//
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1
// 示例 1：
//
// 输入：nums = [1, 7, 3, 6, 5, 6]
// 输出：3
// 解释：
// 中心下标是 3 。
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
// 示例 2：
//
// 输入：nums = [1, 2, 3]
// 输出：-1
// 解释：
// 数组中不存在满足此条件的中心下标。
// 示例 3：
//
// 输入：nums = [2, 1, -1]
// 输出：0
// 解释：
// 中心下标是 0 。
// 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
// 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
func pivotIndex(nums []int) int {

	return 0
}

// 思路
// 1. total, sum(0,i), num[i], sum(i+1, n)： 寻找 sum(0,i-1)+num[i]+sum(i+1,n) = total，如果存在返回i，不存在返回-1
// 2. leftSum=0, rightSum=total-num[0]
// 3.1 leftSum[i] = leftSum + num[i-1]
// 3.2 rightSum[i] = rightSum - num[i]
// if leftSum[i] == rightSum[i] { return i }
func PivotIndexMethod01(nums []int) int {
	// 初始化leftSum和rightSum
	var leftSum, rightSum int
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}

	// 开始迭代数组
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			leftSum += nums[i-1]
		}
		rightSum -= nums[i]

		// 第i位置左侧和右侧sum和相等，返回结果i
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func TestPivotIndexMethod01(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"t2", []int{1, 2, 3}, -1},
		{"t3", []int{2, 1, -1}, 0},
		{"t4", []int{1, -1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PivotIndexMethod01(tt.nums); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
