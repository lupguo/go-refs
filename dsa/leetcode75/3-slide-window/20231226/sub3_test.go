package _0231226

import (
	"testing"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
//
// 输入: s = "pwwkew"
// 输出: 3
// 思路：
//
//	在s上滑动窗口，类似毛毛虫移动
func lengthOfLongestSubstring(s string) int {
	// 记录当前窗口值
	curWin := make(map[byte]int)
	var maxLen int
	var left, right int

	// 开始窗口滑动
	for right < len(s) {
		curWin[s[right]]++
		right++

		// 右字符不在窗口内，更新窗口内容以及最大不重复子串长度
		if isGoodWin(curWin) {
			maxLen = max(maxLen, right-left)
			continue
		}

		// 有重复值，需要左边框右移，检测是否满足不重复子串条件
		for left < right && !isGoodWin(curWin) {
			curWin[s[left]]--
			left++
		}
	}

	return maxLen
}

func isGoodWin(win map[byte]int) bool {
	for _, n := range win {
		if n > 1 {
			return false
		}
	}

	return true
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
