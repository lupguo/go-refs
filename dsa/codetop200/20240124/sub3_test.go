package _0240124

import (
	"testing"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 思路：滑动窗口 , left, right两个指针，类似毛毛虫
// 注意错误:
//
//	a. 没有考虑窗口暂存问题，要用win := make(map[byte]bool)存储
//	b. 联合判断，越界条件放前面，right < n && win[s[right]] == false
//	c. 同理窗口缩小部分，重置代码win[c] = false 放在left++前面，避免切片溢出
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	left, right := 0, 0
	win := make(map[byte]bool)
	maxLen := 0

	for left < n {
		// 窗口扩大
		for right < n && win[s[right]] == false {
			win[s[right]] = true
			right++
			maxLen = max(maxLen, right-left)
			continue
		}

		// 窗口缩小
		win[s[left]] = false
		left++
	}

	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"t1", "abb", 2},
		{"t2", "bbbbb", 1},
		{"t3", "pwwkew", 3},
		{"t4", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)

			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
