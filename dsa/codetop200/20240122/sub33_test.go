package _0240122

import (
	"testing"
)

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 时间复杂度为 O(log n) 的算法解

// 思路: 二分算法，二分变种
//  1. 思考如何折半淘淘掉一半数据，即中间点nums[mid]，识别左侧或右侧是否增量区间，以及target值是否在该范围内，对应的移动left,right游标
//
// 耗时: 2.5h小时
// 错误:
//  1. mid = (right-left+1)/2; 应该是 mid = left + (right-left+1)/2
//  2. 循环条件判断问题: mid >= 0 应该是 mid >= 0 && left <= right（数组越界）
//  3. 条件没有穷尽：nums[mid]==target; 其他情况： nums[left] < nums[mid], nums[left] > nums[mid], 以及default状态(这里遗漏了)
func search(nums []int, target int) int {
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

func TestSearchRound(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"t1", []int{1}, 1, 0},
		{"t2", []int{1, 2}, 1, 0},
		{"t3", []int{1, 2}, 2, 1},
		{"t4", []int{1, 2}, 3, -1},
		{"t5", []int{1, 2, 3}, 3, 2},
		{"t6", []int{2, 1}, 1, 1},
		{"t7", []int{2, 1}, 2, 0},
		{"t8", []int{3, 1, 2}, 3, 0},
		{"t9", []int{3, 4, 1, 2}, 2, 3},
		{"t10", []int{3, 4, 1, 2}, 1, 2},
		{"t11", []int{3, 4, 1, 2}, 5, -1},
		{"t12", []int{4, 5, 1, 3}, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}

// 二分查找思路
//  1. left,mid,right，分成[left,mid), [mid,right] 两部分
//  2. 循环条件 left <= right，且target需要在查询范围内
//  3. 中间值mid := left+(right-left+1)/2
//  4. 和中间值比较，重新设置left,right值，减半处理
//
// 错误点:
//  1. 不在判断范围外 target <= nums[left]
//  2. right = mid，right应该是mid-1
func binarySearch(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right { // 循环满足基础条件
		if target < nums[left] || target > nums[right] { // 不在范围内
			return -1
		}
		// 中间值
		mid := left + (right-left+1)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid
		case nums[mid] > target:
			right = mid - 1
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"t1", []int{1}, 1, 0},
		{"t2", []int{1}, 3, -1},
		{"t3", []int{1, 2}, 1, 0},
		{"t4", []int{1, 2}, 2, 1},
		{"t5", []int{1, 2, 3}, 2, 1},
		{"t6", []int{1, 2, 3}, 3, 2},
		{"t7", []int{1, 2, 3}, 3, 2},
		{"t8", []int{1, 2, 3}, 3, 2},
		{"t9", []int{1, 3, 4, 5, 10}, 5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})

	}
}
