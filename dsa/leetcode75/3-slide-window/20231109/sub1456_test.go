package _0231109

import (
	"strings"
	"testing"
)

// 给你字符串 s 和整数 k 。
// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
// 英文中的 元音字母 为（a, e, i, o, u）。
//
// 示例 1：
//
// 输入：s = "abciiidef", k = 3
// 输出：3
// 解释：子字符串 "iii" 包含 3 个元音字母。
// 示例 2：
//
// 输入：s = "aeiou", k = 2
// 输出：2
// 解释：任意长度为 2 的子字符串都包含 2 个元音字母。
// 示例 3：
//
// 输入：s = "leetcode", k = 3
// 输出：2
// 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
// 示例 4：
//
// 输入：s = "rhythms", k = 4
// 输出：0
// 解释：字符串 s 中不含任何元音字母。
// 示例 5：
//
// 输入：s = "tryhard", k = 4
// 输出：1
func maxVowels(s string, k int) int {

	return 0
}

// 滑动窗口代码优化
func Method2MaxVowels(s string, k int) int {
	vowels := "aeiou"
	n := len(s)
	var curCnt, maxCnt int
	for i := 0; i < n; i++ {
		// 第i位置是否有元音字符
		if strings.Index(vowels, string(s[i])) >= 0 {
			curCnt++
		}

		// 当i移动到k位置，这需要检视i-k位置，是否为有元音字符，有则移除
		if i >= k && strings.Index(vowels, string(s[i-k])) >= 0 {
			curCnt--
		}
		maxCnt = max(maxCnt, curCnt)
	}

	return maxCnt
}

// 滑动窗口
// 1. k个字符子串，指针i从 0<= i < n-k+1，位置移动，假定cntk[i]表示从指针位置i到i+k+1位置的拥有元音数
// 那么cnk[i+1] = cntk[i] + s[i+k-1]是否为元音(before=1 or 0) - s[i]是否为元音(after=1 or 0)
func Method1MaxVowels(s string, k int) int {
	n := len(s)
	if k > n {
		return 0
	}

	// 元音字符map初始化
	vowMap := make(map[byte]bool)
	for _, c := range "aeiou" {
		vowMap[byte(c)] = true
	}

	// 先统计cntk[0]，默认第一个k串就是最大值
	var maxVowelCnt int
	for i := 0; i < k; i++ {
		if vowMap[s[i]] { // i位置为元音
			maxVowelCnt++
		}
	}

	// 开始从cntk[1]，利用滑动窗口公式迭代
	beforeVowelCnt := maxVowelCnt
	var curKVowelCnt int
	for i := 1; i < n-k+1; i++ {
		// s[i-1] 前一个字符是否元音
		var before, addition int
		if vowMap[s[i-1]] {
			before = 1
		}

		// s[i+k-1] 当前字符串是否为元音
		if vowMap[s[i+k-1]] {
			addition = 1
		}

		// 利用公式迭代出每个
		curKVowelCnt = beforeVowelCnt + addition - before
		if curKVowelCnt > maxVowelCnt {
			maxVowelCnt = curKVowelCnt
		}
		beforeVowelCnt = curKVowelCnt
	}

	return maxVowelCnt
}

func TestMethod1MaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"t1", "abciiidef", 3, 3},
		{"t2", "aeiou", 2, 2},
		{"t3", "leetcode", 3, 2},
		{"t4", "rhythms", 4, 0},
		{"t5", "tryhard", 4, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method2MaxVowels(tt.s, tt.k); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
