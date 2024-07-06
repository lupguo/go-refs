package _0240306

import (
	"testing"
)

// 有序数组旋转后的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 思路: 二分 O(LogN)
//
//	4,5,6,7,0,1,2 -> low, mid, hi
//	1. 去数组中值v=nums[n/2], 检测hi位置元素是否 > mid，是的则从该半部分折半查找
//
// 问题:
//  1. 需要充分利用一侧自增特性，左折半查找
//  2. 返回结果需要检测是否为-1
//  3. 考虑到high边界问题
func search(nums []int, target int) int {
	// base case
	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	low, mid, high := 0, n/2, n-1

	if target == nums[mid] {
		return mid
	}
	if nums[low] < nums[mid] { // 左侧自增
		if nums[low] <= target && target < nums[mid] {
			return searchIndex(nums[low:mid], target, low)
		} else {
			return searchIndex(nums[mid:], target, mid)
		}
	} else { // 右侧自增
		if target > nums[mid] && target <= nums[high] {
			return searchIndex(nums[mid:], target, mid)
		} else {
			return searchIndex(nums[low:mid], target, low)
		}
	}
}

func searchIndex(nums []int, target int, startIndex int) int {
	if foundIndex := search(nums, target); foundIndex != -1 {
		return foundIndex + startIndex
	} else {
		return -1
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"t1", []int{1, 2}, 1, 0},
		{"t2", []int{1, 2}, 2, 1},
		{"t3", []int{1, 2}, 3, -1},
		{"t4", []int{3, 1, 2}, 3, 0},
		{"t5", []int{3, 1, 2}, 1, 1},
		{"t6", []int{3, 1, 2}, 2, 2},
		{"t7", []int{3, 4, 1, 2}, 3, 0},
		{"t8", []int{3, 4, 1, 2}, 4, 1},
		{"t9", []int{3, 4, 1, 2}, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search2(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})

	}
}

func search2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
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
