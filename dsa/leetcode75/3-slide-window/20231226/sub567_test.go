package _0231226

import (
	"testing"
)

// 总结：
//  1. 要先看清题目，具体是字符串的排列（组合），还是子串（有序）问题，两者采取的方法不同
//  2. 找子串问题，双指针就可以解决; 排列问题，需要滑动窗口（这里是固定窗口大小）
//  3. 滑动大致思路
//     a. 先确认窗口要求needWin字符计数
//     b. 创建一个curWin当前窗口，以及左右边框指针left, right
//     c. 尝试right在合理范围值内（right < s2.len) ，进行窗口扩张，进行right++移动
//     d. 窗口扩张后检测是否符合匹配基本条件，比如固定窗口（s2[left:right] == s1.len) 才停止右边框右滑
//     f. 右边框确认后，检测是否有符合窗口匹配：isMeetWin(curWin, needWin)，符合返回，不符合，左边框右移
//     g. 因为这里只涉及固定窗口大小，所以左边框移动无需for循环，直接curWin[v]--; left++右移即可（有的题目需要for循环，类似毛毛虫移动左边框）
//  4. 细节问题
//     a. 窗口过小，右边框要持续扩张
//     b. 窗口都是在s2原始串上移动
//
// 看清楚题目（排列）：s1 的排列之一是 s2 的 子串  -> 有组合问题，窗口
// 检测字符串是否为另一个的子串
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
// 思路: 在s2上窗口滑动，检测是和s1的排列匹配
func checkInclusion(s1 string, s2 string) bool {
	// 异常边界条件
	if len(s2) < len(s1) {
		return false
	}

	// 待比较s1子串窗口
	needWin := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		needWin[s1[i]]++
	}

	// 在s2上窗口滑动，检测是和s1的排列匹配
	curWin := make(map[byte]int)
	var left, right int
	for right < len(s2) {
		// 窗口右边框，右移
		curWin[s2[right]]++
		right++

		// 窗口过小，继续右移
		if right-left < len(s1) {
			continue
		}

		// 窗口一致，检查窗口是否满足条件，满足则直接返回
		if isMeetWindow(curWin, needWin) {
			return true
		}

		// 窗口左边框，右移
		curWin[s2[left]]--
		left++
	}

	return false
}

// 判断s1是否s2的子串(有严格排序，双指针解决）
func isSubStr(s1, s2 string) bool {
	// 原始串i指针不断右移
	for i := 0; i < len(s2); i++ {
		j := 0
		// 子串、原始串是否首字符相等
		if s2[i] != s1[j] {
			continue
		}

		// 子串满足匹配，则原始串、子串双指针移动
		for j < len(s1) && i < len(s2) && s2[i] == s1[j] {
			i++
			j++
		}

		// 看子串是否迭代完成
		if len(s1) == j {
			return true
		}
	}

	return false
}

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"t1", "a", "ab", true},
		{"t2", "b", "ab", true},
		{"t3", "ba", "ab", true},
		{"t4", "ab", "bab", true},
		{"t5", "aba", "bab", false},
		{"t6", "ab", "eidbaooo", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
