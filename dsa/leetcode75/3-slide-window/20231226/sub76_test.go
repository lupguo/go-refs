package _0231226

import (
	"testing"
)

// 源字符串s最小覆盖子串t
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
func minWindow(s string, t string) string {
	// 参数定义
	var minWinStr string            // 最小窗口字符串
	curWinCnt := make(map[byte]int) // 当前窗口

	// 子串目标窗口要求
	needWinCnt := make(map[byte]int)
	for i, _ := range t {
		needWinCnt[t[i]]++
	}

	// 开始右移右边框寻找可行解
	left, right := 0, 0  // 左右窗口边界
	for right < len(s) { // 不超过右边界
		curWinCnt[s[right]]++                     // 当前串窗口计数更新
		right++                                   // 右边框右移
		if !isMeetWindow(curWinCnt, needWinCnt) { // 移动后仍然不满足目标窗口需求，右边框继续右滑动，
			continue
		}

		// curWin满足子串可行解，检测是否需要更新minWin
		minWinStr = swapMinWin(s[left:right], minWinStr)

		// 左边界右滑动，从可行解里面找最优解
		for left < right {
			curWinCnt[s[left]]--                      // 左边框计数更新
			left++                                    // 左边框右移
			if !isMeetWindow(curWinCnt, needWinCnt) { // 移动后，不满足子串条件了，结束本轮左窗口的右滑动
				break
			}

			// curWin满足子串可行解，尝试更新最小窗口
			minWinStr = swapMinWin(s[left:right], minWinStr)
		}

	}

	return minWinStr
}

// 交换最小窗口
func swapMinWin(curWin string, minWin string) string {
	if minWin == "" {
		return curWin
	}
	if len(curWin) < len(minWin) {
		return curWin
	}
	return minWin
}

// 检测curWin字符是否满足当前t的结果
func isMeetWindow(curWin map[byte]int, needWin map[byte]int) bool {
	if len(curWin) < len(needWin) {
		return false
	}

	// 检测是否mc中还有剩余字符，若存在则表示不满足
	for c, cnt := range needWin {
		if curWin[c] < cnt {
			return false
		}
	}

	return true
}

func TestMinWin(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"t1", "a", "A", ""},
		{"t2", "aA", "A", "A"},
		{"t3", "AaA", "Aa", "Aa"},
		{"t4", "AaA", "aA", "Aa"},
		{"t5", "AXYBA", "AB", "BA"},
		{"t6", "AXYBBA", "AB", "BA"},
		{"t7", "AXYBBA", "ABB", "BBA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
