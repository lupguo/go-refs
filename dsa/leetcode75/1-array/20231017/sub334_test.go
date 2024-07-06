package _0231017

import (
	"reflect"
	"testing"
)

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
func increasingTriplet(nums []int) bool {

	return method2increasingTriplet(nums)
}

// todo: 使用贪心算法求解这个三元组问题
func method3increasingTriplet(nums []int) bool {

	return false
}

// 通过分区间，原题等价于从nums[0:j] 找最小值 、从nums[j:len-1]中找最大值
func method2increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// 初始化一个数组，存储maxRight - 这里是用了空间换时间，不这样处理，在检测num[j]时候，会出现O(N^2)时间复杂度
	rightMaxNums := make([]int, n)
	rightMaxNums[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMaxNums[i] = max(nums[i], rightMaxNums[i+1])
	}

	// ps. 因为左侧符合递增规则，无需额外数组存储minLeftNums，可以直接借助一个变量minLeft即可
	leftMinNums := make([]int, n)
	leftMinNums[0] = nums[0]
	for i := 1; i <= n-2; i++ {
		leftMinNums[i] = min(leftMinNums[i-1], nums[i])
	}

	// 检测符合条件的num[j]
	for j := 1; j <= n-2; j++ {
		// 如果左侧区间num[0:j]的最小值都比num[j]还大，则表示num[i] < num[j]肯定不满足，则向后寻找中间值num[j]
		if leftMinNums[j] >= nums[j] {
			continue
		}

		// 如果右侧区间num[j+1:]的最大值没有超过num[j]，则表示num[j] < num[k]肯定无法满足，继续向后寻找中间值num[j]
		if nums[j] >= rightMaxNums[j+1] {
			continue
		}

		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 输入：nums = [2,1,5,0,4,6]
// 输出：true
// 解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
// 暴力算法
func method1increasingTriplet(nums []int) bool {
	// 依次i,j,k检测是否存在nums[i] < nums[j] < nums[k]
	length := len(nums)
	// i的检测范围[0,len-3]
	for i := 0; i <= length-3; i++ {

		// j的检测范围[0,len-2]
		for j := i + 1; j <= length-2; j++ {

			if nums[j] <= nums[i] {
				continue
			}

			for k := j + 1; k <= length-1; k++ {

				if nums[k] <= nums[j] {
					continue
				}

				return true
			}
		}
	}

	return false
}

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"t1", []int{1, 2, 3, 4}, true},
		{"t2", []int{5, 2, 3, 4, 1}, true},
		{"t3", []int{5, 3, 1, 2, 2}, false},
		{"t4", []int{2, 1, 5, 0, 4, 6}, true},
		{"t5", []int{1, 5, 0, 4, 1, 3}, true},
		{"t6", []int{5, 1, 6}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := method2increasingTriplet(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})

	}
}
