package _0240223

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
//
// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
//
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
//
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
//
// 你所设计的解决方案必须只使用常量级的额外空间。
// 思路: 非递减顺序排列, 有序left, right 指针夹逼
func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	var sumTwo int
	for left < right {
		sumTwo = numbers[left] + numbers[right]

		// sum和小了，left 右移增大sum和，继续比较
		for left < right && sumTwo < target {
			left++
			sumTwo = numbers[left] + numbers[right]
		}

		// sum和大了，right 左移
		for left < right && sumTwo > target {
			right--
			sumTwo = numbers[left] + numbers[right]
		}

		// 相等
		if sumTwo == target {
			return []int{left, right}
		}
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"t1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"t2", []int{3, 2, 4}, 6, nil},
		{"t3", []int{3, 3}, 6, []int{0, 1}},
		{"t4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17, []int{7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
