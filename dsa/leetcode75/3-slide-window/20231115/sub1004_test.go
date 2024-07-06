package _0231115

import (
	"testing"
)

// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
//
// 示例 1：
//
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
// 示例 2：
//
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
func longestOnes(nums []int, k int) int {

	return 0
}

// 连续k个0，最大区间，即为求解值
// 最大个数 -> 滑动窗口解题思路
// 题目转换成，寻找数组中连续区间(left, right)内，zeroCnt包含不超过k的最长数组的长度(right-left+1)
func MethodLongestOnes1(nums []int, k int) int {
	n := len(nums)

	// left 记录数组最左侧为0
	var left int

	// 连续数组区间内，零值计算
	var zeroCnt int

	// 最大长度
	var maxLen int

	// right 从初始位置开始迭代
	for right := 0; right < n; right++ {

		// right 位置遇到0
		if nums[right] == 0 {

			// 零值计算增加
			zeroCnt++

			// 若计数未超过k阈值，计算最大长度，right继续右移
			if zeroCnt <= k {
				maxLen = max(maxLen, right-left+1)
				continue
			}

			// 若计数超过了k阈值,表示区间内的零过多了，左侧的left边界需要不断向右移动，寻找到一个零值的下个位置，同时计数器减一
			for ; left <= right; left++ {
				if nums[left] == 0 {
					left++
					zeroCnt--
					break
				}
			}

		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func TestMethodLongestOnes1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"t1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"t2", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{"t3", []int{0, 0, 1, 1, 1, 0, 0}, 0, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MethodLongestOnes1(tt.nums, tt.k); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
