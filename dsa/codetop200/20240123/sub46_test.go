package _0240123

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/permutations/
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 全排列问题：回溯算法
//   - 根->循环选择列表做出选择->做出子选择..->节点，
//   - 遍历到节点过程就是答案
func permute(nums []int) [][]int {
	// 基础条件&异常case
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	for _, num := range nums {
		// 做出选择
		// var ans []int
		// ans = append(ans, num)

		// 剩余选择
		leftNums := removeElement(nums, num)
		subPermutes := permute(leftNums)

		// 针对每个子nums可以返回的全排列，拼接上当前选择的数据
		for i, _ := range subPermutes {
			subPermutes[i] = append(subPermutes[i], num)
		}

		ret = append(ret, subPermutes...)
	}

	return ret
}

func removeElement[T comparable](slice []T, elem T) []T {
	var result []T
	for _, v := range slice {
		if v != elem {
			result = append(result, v)
		}
	}
	return result
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		elem      int
		wantSlice []int
	}{
		{"t1", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{"t2", []int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
		{"t3", []int{1, 1, 1, 1, 1}, 1, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.slice, tt.elem)

			if len(got) != len(tt.wantSlice) {
				t.Errorf("got %v, but want %v", got, tt.wantSlice)
			}

			for i := range got {
				if got[i] != tt.wantSlice[i] {
					t.Errorf("got %v, but want %v", got, tt.wantSlice)
					break
				}
			}
		})
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"t1", []int{1}, [][]int{{1}}},
		{"t2", []int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{"t3", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}

			for _, ans := range got {
				found := false
				for _, want := range tt.want {
					if reflect.DeepEqual(ans, want) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("got %v, but want %v", got, tt.want)
				}
			}
		})
	}
}
