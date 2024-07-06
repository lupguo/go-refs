package _0231016

import (
	"reflect"
	"testing"
)

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
func productExceptSelf(nums []int) []int {

	return method1ProductExceptSelf(nums)
}

// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
func method1ProductExceptSelf(nums []int) []int {
	// 左侧
	leftMultVal := make([]int, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		if i == 0 {
			leftMultVal[0] = 1
		} else {
			leftMultVal[i] = leftMultVal[i-1] * nums[i-1]
		}
	}

	// 右侧元素乘积
	rightMultVal := make([]int, len(nums))
	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			rightMultVal[j] = 1
		} else {
			rightMultVal[j] = rightMultVal[j+1] * nums[j+1]
		}
	}

	// 左侧x右侧，为指定元素答案
	ans := make([]int, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		ans[i] = leftMultVal[i] * rightMultVal[i]
	}

	return ans
}

func TestMethod1ProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"t2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := method1ProductExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
