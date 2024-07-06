package _0231116

import (
	"testing"
)

// 给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
//
// 如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
// 示例 1：
//
// 输入：arr = [1,2,2,1,1,3]
// 输出：true
// 解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
// 示例 2：
//
// 输入：arr = [1,2]
// 输出：false
// 示例 3：
//
// 输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
// 输出：true
func uniqueOccurrences(arr []int) bool {

	return false
}

// 思路:
// 1. 数组 转map map[int]int，值为数的计数次数
// 2. map中每个数字的计数是否独一无二，新创建一个 map[int]bool, key 为计数次数，如果计数次数存在，返回false, 否则最终返回true
func UniqueOccurrences(arr []int) bool {
	numCntMap := make(map[int]int) // 数字计数次数map
	exist := make(map[int]bool)    // 计数次数是否存在
	for _, v := range arr {
		numCntMap[v]++
	}

	for _, cnt := range numCntMap {
		// 数字计数的次数是否是独一无二的
		if exist[cnt] {
			return false
		}

		exist[cnt] = true
	}

	return true
}

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"t1", []int{1, 2, 2, 1, 1, 3}, true},
		{"t2", []int{1, 2}, false},
		{"t3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueOccurrences(tt.arr); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
