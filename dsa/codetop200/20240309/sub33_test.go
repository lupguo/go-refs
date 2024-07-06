package _0240309

import (
	"testing"
)

const NotFound = -1

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
//
// 针对有序数组旋转后，得到nums，从该nums查询target获取对应的下标，找不到返回-1
// 必须设计一个时间复杂度为 O(log n) 的算法解决此问题
// [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 思路:
//   - 二分查找，从有明确特征到半边(升序）查找，low, mid, high = 0, n/2, n-1
//     例如查找6, 判断nums[mid]=7, 左边nums[low] < nums[mid] 是[4,..7]是递增区间, 且 target 在区间内，则在该部分使用search二分查找（更新 high=mid)
//     例如查找2, 判断nums[mid]=7, 左边nums[low] < nums[mid] 是[4,..7]是递增区间, 但 target 不在区间内，则在另外半边使用search算法（更新 low=mid+1)
func search(nums []int, target int) int {
	// base case
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 二分查找
	low, mid, high := 0, n/2, n-1
	switch {
	case target == nums[mid]:
		return mid
	case nums[low] < nums[mid]: // 左增量
		if nums[low] <= target && target <= nums[mid] { // target 在左增区间内
			return low + search(nums[low:mid], target)
		} else {
			return mid + 1 + search(nums[mid+1:], target)
		}
	default: // 右增量区间
		if nums[mid] <= target && target <= nums[high] { // target 在右增量区间
			return mid + 1 + search(nums[mid+1:high], target)
		} else {
			return low + search(nums[low:mid], target)
		}
	}
}

func wrap(index int, fn func(nums []int, target int) int) {

}

// 迭代法
func search2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for mid := (right - left + 1) / 2; mid >= 0 && left <= right; mid = left + (right-left+1)/2 {
		switch {
		case nums[mid] == target:
			return mid
		case nums[left] < nums[mid]: // 左边是自增的
			if nums[left] <= target && target < nums[mid] { // 目标值在左侧自增区间内
				right = mid - 1
			} else {
				left = mid + 1
			}
		case nums[left] > nums[mid]: //  左侧出现了阶梯断层，考虑右侧自增区间是否存在target值
			if nums[mid] < target && target <= nums[right] { // 目标值在右侧自增区间内
				left = mid + 1
			} else {
				right = mid - 1
			}
		default:
			return -1
		}
	}

	return -1
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"t1", []int{1}, 1, 0},
		{"t2", []int{1}, 2, -1},
		{"t3", []int{1, 2}, 2, 1},
		{"t4", []int{2, 1}, 2, 0},
		{"t5", []int{2, 1}, 3, -1},
		{"t6", []int{1, 3, 5}, 5, 2},
		{"t7", []int{7, 8, 1, 2, 3, 4, 5, 6}, 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
