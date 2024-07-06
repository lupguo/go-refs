package _0240122

import (
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
// 思路: 滑动窗口解法
//  1. left, right: 窗口左右指针
//  2. win 窗口map
//  3. maxWinSize 最大打窗口大小
//     先right移动，right++扩大窗口，知道遇到和窗口内重复字符 -> 左边界left缩小窗口left++，直到解除重复字符 -> 重新右扩窗口
//
// 问题:
//  1. 窗口计算应该是在参数迭代之前完成: maxWinSize = max(maxWinSize, right-left+1)
//
// 耗时: 30min
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	left, right := 0, 0
	win := make(map[byte]int)
	maxWinSize := 0

	// 迭代字符串
	for left <= right && right < n {
		// 右移right窗口
		if win[s[right]] == 0 { // 增大窗口
			win[s[right]]++
			maxWinSize = max(maxWinSize, right-left+1)
			right++
			continue
		}

		// 遇到窗口中重复字符，缩小窗口
		for left < right {
			if s[left] != s[right] {
				win[s[left]]--
				left++
				continue
			}
			win[s[left]]--
			left++
			break
		}

	}

	return maxWinSize
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"t1", "ab", 2},
		{"t2", "aab", 2},
		{"t3", "aaba", 2},
		{"t4", "acaba", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
