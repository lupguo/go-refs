package _0240306

import (
	"testing"
)

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数
// https://leetcode.cn/problems/first-missing-positive/description/
// 思路:
//  1. 申请一个大小为nums.length长度的map（O(N)空间复杂度），将nums元素放入map中
//  2. 将i从1迭代到n，检测i是否在map中，不在则返回i，否则返回i+1
func firstMissingPositive(nums []int) int {
	// 载入map - O(N)空间复杂度 - 非常数级
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	// 遍历检测
	i := 1
	for ; i <= len(nums); i++ {
		if m[i] == false {
			return i
		}
	}

	return i
}

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1}, 2},
		{"t2", []int{0}, 1},
		{"t3", []int{0, 1}, 2},
		{"t4", []int{0, 1, 3}, 2},
		{"t5", []int{0, 1, 2, 7, 8}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstMissingPositive(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
