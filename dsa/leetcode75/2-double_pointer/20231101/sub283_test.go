package _0231101

import (
	"testing"
)

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例 1:
//
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 示例 2:
//
// 输入: nums = [0]
// 输出: [0]

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroes(nums []int) {
	nums = method1moveZeroes(nums)
}

// 【不复制数组】 - 两次遍历
// 过O(N)时间复杂度遍历数组，将所有非0依次更新到数组最左侧，剩余的[j,length]重置成0
func method2moveZeroes(nums []int) []int {
	length := len(nums)

	// 过O(N)时间复杂度遍历数组，将所有非0依次更新到数组最左侧
	var j int
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// 剩余的[j,length]重置成0
	for i := j; i < length; i++ {
		nums[i] = 0
	}

	return nums
}

// 【不复制数组】 - 一次遍历，参考快排思想
// 参考快排思想，将不等于0的nums[i]和等于0的nums[j]进行值的替换
func method3moveZeroes(nums []int) []int {
	var j int
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}
	return nums
}

// 1. 新创建一个数组，长度为元素数组长度
// 2. 迭代数组：(初始尾部tail指针位置)
//   - 元素i的值为0 -> 放到tail指针位置，指针向前移动一格
//   - 元素i的不为0 -> 放到head指针位置，指针向后移动一格
func method1moveZeroes(nums []int) []int {
	// 新创建数组
	length := len(nums)
	newArrs := make([]int, length)
	head, tail := 0, length-1
	for _, num := range nums {
		if num == 0 {
			newArrs[tail] = num
			tail--
		} else {
			newArrs[head] = num
			head++
		}
	}

	return newArrs
}

func TestMethod1moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{0, 1, 0, 2}, []int{1, 2, 0, 0}},
		{"t2", []int{0}, []int{0}},
		{"t3", []int{1, 0, 0}, []int{1, 0, 0}},
		{"t4", []int{1, 0, 0, 2}, []int{1, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := method3moveZeroes(tt.nums)
			for i, _ := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, but want %v, got[i]=%v, nums[i]=%v", got, tt.want, got[i], tt.want[i])
				}
			}
		})
	}
}
