package _0231116

import (
	"testing"
)

// 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：
//
// answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
// answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
// 注意：列表中的整数可以按 任意 顺序返回。

// 示例 1：
//
// 输入：nums1 = [1,2,3], nums2 = [2,4,6]
// 输出：[[1,3],[4,6]]
// 解释：
// 对于 nums1 ，nums1[1] = 2 出现在 nums2 中下标 0 处，然而 nums1[0] = 1 和 nums1[2] = 3 没有出现在 nums2 中。因此，answer[0] = [1,3]。
// 对于 nums2 ，nums2[0] = 2 出现在 nums1 中下标 1 处，然而 nums2[1] = 4 和 nums2[2] = 6 没有出现在 nums2 中。因此，answer[1] = [4,6]。
// 示例 2：
//
// 输入：nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// 输出：[[3],[]]
// 解释：
// 对于 nums1 ，nums1[2] 和 nums1[3] 没有出现在 nums2 中。由于 nums1[2] == nums1[3] ，二者的值只需要在 answer[0] 中出现一次，故 answer[0] = [3]。
// nums2 中的每个整数都在 nums1 中出现，因此，answer[1] = [] 。
func findDifference(nums1 []int, nums2 []int) [][]int {

	return nil
}

// 寻找数组差异
// 1. nums1 转map => m1, map[int]bool, num2 转map =>m 2
// 2. 初始化二维数组 answer
// 3.1 answer[0]: 遍历 nums1, num not int m2, append to answer[0]
// 3.2 answer[1]: 遍历 nums2, num not int m1, append to answer[1]
func FindDifferenceMethod01(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]bool), make(map[int]bool)

	// 初始化m1,m2
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}

	// 初始化一个结果数组
	answer := make([][]int, 2)

	// 填充answer: 在nums1中检测元素是否不在nums2中出现过，满足条件追加到数组中
	// 要求新添加的元素不能一致，所以这里需要重新申请两个map做几路
	mm1, mm2 := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		if !m2[v] && mm1[v] == false {
			answer[0] = append(answer[0], v)
			mm1[v] = true
		}
	}

	for _, v := range nums2 {
		if !m1[v] && mm2[v] == false {
			answer[1] = append(answer[1], v)
			mm2[v] = true
		}
	}

	return answer
}

func TestFindDifferenceMethod01(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{"t1", []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{"t2", []int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDifferenceMethod01(tt.nums1, tt.nums2)

			for i, ans := range got {
				for j, _ := range ans {
					if got[i][j] != tt.want[i][j] {
						t.Fatalf("got %v, but want %v", got, tt.want)
					}
				}
			}
		})
	}
}
