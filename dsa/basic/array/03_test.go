package array

import (
	"testing"
)

// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Test03(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	t.Logf("%v", rotate(nums, k))

}

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
func rotate(nums []int, k int) []int {
	size := len(nums)
	if k < 0 || k > size {
		return nums
	}

	// 直接把[k:]放left,[:size-k]放right
	left := nums[:size-k]
	right := nums[size-k:]

	onums := make([]int, size)
	for i := 0; i < size; i++ {
		if i < k {
			onums[i] = right[i]
		} else {
			onums[i] = left[i-k]
		}
	}

	copy(nums, onums)
	// nums = onums

	return nums
}

// newNums := make([]int, size)
// newNums = append(newNums, left...)
// newNums = append(newNums, right...)
//
// nums = newNums
