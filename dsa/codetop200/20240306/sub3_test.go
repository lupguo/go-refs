package _0240306

import (
	"testing"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// s = "abcabcbb" return=3
//
//	"pwwkew" return=3
//
// 思路: 滑动窗口
//  1. win窗口[i,j], 满足条件(窗口内无重复字符), j向右扩张(win增大,更新maxWinSize)，反之i向右扩张(win缩小)
//     注意: j能否扩展，依赖于win是否有重复元素(map判断),无重复则j可以扩展
//     i右移，窗口缩小也需要检测是否满足条件，
func lengthOfLongestSubstring(s string) int {

	win := make(map[byte]bool) // 是否已在窗口中
	var i, j int
	var maxWinSize int
	for i < len(s) && j < len(s) {

		// 将s[j]放入窗口，检测是否符合条件
		for j < len(s) && win[s[j]] == false { // 窗口扩展
			win[s[j]] = true
			maxWinSize = max(j-i+1, maxWinSize)
			j++
		}

		// 检测j是否已经到边界了，不用左侧再收缩了
		if j == len(s) {
			return maxWinSize
		}

		// 窗口内不满足条件，窗口左边缩小
		for i < j && win[s[j]] == true { // 只要窗口里面还有j索引字符，则继续缩小
			win[s[i]] = false
			i++
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
		{"t1", "a", 1},
		{"t2", "aa", 1},
		{"t3", "aab", 2},
		{"t4", "abacb", 3},
		{"t5", "adbbac", 3},
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
