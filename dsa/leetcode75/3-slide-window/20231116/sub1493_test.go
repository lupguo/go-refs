package _0231116

import (
	"testing"
)

// 给你一个二进制数组 nums ，你需要从中删掉一个元素（必须删除一个元素），返回 删掉一个元素以后全为 1 的最长子数组
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
//
// case 1
// 输入：nums = [1,1,0,1]
// 输出：3
// 解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
//
// case 2
// 输入：nums = [0,1,1,1,0,1,1,0,1]
// 输出：5
// 解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
//
// case 3
// 输入：nums = [1,1,1]
// 输出：2
// 解释：你必须要删除一个元素。
func longestSubarray(nums []int) int {

	return 0
}

// 从左向右移动，right探路，同时记录zeroCnt数量（超过 or 未超过）
// 未超过: 左侧left不用动，统计并更新maxLen长度
// 超过: 左侧left被迫移动到下个0的后一位，使zeroCnt不超过阈值
func LongestSubarrayMethod01(nums []int) int {
	n := len(nums)

	// 可删除0值的最大长度，这里是k=1
	k := 1

	// 连续子数组左边界，当子数组内0超过了1个则必须右移
	var left int

	// 来个计数值0
	var zeroCnt int

	// maxLen为删除k个0值最大子数组长度
	var maxLen int

	// right为连续子数组右边界，不断右移
	for right := 0; right < n; right++ {

		// 右移过程遇到0
		if nums[right] == 0 {
			zeroCnt++

			// 如果没有超过阈值
			if zeroCnt <= k {
				// 最大长度为左右边界内元素个数 - zeroCnt
				maxLen = max(maxLen, right-left+1-zeroCnt)
				continue
			}

			// 如果超过了阈值，左侧left被迫右移到下个0位置
			for ; left < right; left++ {
				if nums[left] == 0 {
					left++
					zeroCnt--
					break
				}
			}
		}

		// 右移过程，如果zeroCnt为0，则需要扣减k，否则扣减zeroCnt
		if zeroCnt == 0 {
			maxLen = max(maxLen, right-left+1-k)
		} else {
			maxLen = max(maxLen, right-left+1-zeroCnt)
		}
	}

	return maxLen
}

func TestLongestSubarrayMethod01(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1, 1, 0, 1}, 3},
		{"t2", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{"t3", []int{1, 1, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubarrayMethod01(tt.nums); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
