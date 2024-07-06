package _0240308

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/3sum/
// 返回所有和为 0 且不重复的三元组
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 思路:
//  1. 先排序，O(NLogN)
//  2. 从[-1,-1,0,1,2] 中找和为0的，遍历nums, 从剩余的列表nums[i:]中，寻找twoSum(nums[i:],-nums[i]) []int 的结果，和nums[i]组合返回
//  3. 两数之和使用双指针相向而行靠拢，因为答案要不重复的解，所以需要注意去重（法1：考虑使用一个exist map解决, 法2：因为有序，使用一个last变量更替即可）
func threeSum(nums []int) [][]int {
	// 基础检测
	if len(nums) < 3 {
		return nil
	}

	// 排序
	sort.Ints(nums)

	// 迭代nums
	var ans [][]int
	last := math.MinInt
	for i, num := range nums {
		// 上个数字已经配对过，则跳过不再复用了
		if last == num {
			continue
		}

		// 从剩余nums[i:]中，寻找两数和为-nums[i]的
		if twoNums := twoSumSortedNums(nums[i+1:], -nums[i]); twoNums != nil {
			for _, two := range twoNums {
				ans = append(ans, []int{num, two[0], two[1]})
			}
		}
		last = num
	}

	return ans
}

// 从有序nums中查找和为target的两个列表
func twoSumSortedNums(nums []int, target int) [][]int {
	var ans [][]int

	// i, j 向中间靠拢
	i, j := 0, len(nums)-1
	last := math.MinInt
	for i < j {
		twoSum := nums[i] + nums[j]
		if twoSum == target { // 确保i不重复
			// 是否i,j 已使用过了
			if last != nums[i] {
				ans = append(ans, []int{nums[i], nums[j]})
				last = nums[i]
			}
			i++
			j--
		} else if twoSum < target {
			i++
		} else {
			j--
		}
	}

	return ans
}

func TestThreeNums(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"t1", []int{-1, -1}, nil},
		{"t2", []int{-1, -1, 2}, [][]int{{-1, -1, 2}}},
		{"t3", []int{-1, -1, 2, 2}, [][]int{{-1, -1, 2}}},
		{"t4", []int{-1, -1, 2, 0}, [][]int{{-1, -1, 2}}},
		{"t5", []int{-1, -1, 2, 0, 1}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"t6", []int{1, 2, -2, -1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}

			// sum和
			for _, nums := range got {
				var sumThree int
				for _, num := range nums {
					sumThree += num
				}
				if sumThree != 0 {
					t.Errorf("sum three nums got %v, nums %v, got %v", sumThree, nums, got)
				}
			}
		})
	}
}
