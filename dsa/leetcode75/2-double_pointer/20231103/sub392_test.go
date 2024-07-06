package _0231103

import (
	"testing"
)

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
func isSubsequence(s string, t string) bool {

	return Method1IsSubSequence(s, t)
}

// 思路: i,j 双指针，分别挂在s 和 t字符串上，要求i,j 在可迭代范围内
// i 表示从t中找了匹配的字符
// j 表示不断向前移动指针
// 不满足条件后(不管是i还是j迭代完后），检测i是不是和原始的s 串长度一致，如果一致则表示找到目标值，返回true
func Method2Method2IsSubSequence(s string, t string) bool {
	// i,j 表示指向字符i,j位置的指针
	var i, j int
	lens, lent := len(s), len(t)

	// 任意i, j 迭代完成后退出
	for i < lens && j < lent {
		if s[i] == t[j] { // 相等i后移
			i++
		}
		j++ // j不断后移
	}

	// 如果子串迭代完成，则表示找到了匹配的子串
	return i == lens
}

// 依次位置0，1...N 从子串中找字符(s[i])，将找到的字符依次迭代从t中寻找(比如在t[j]位置)，如果找到在位置则继续从子串s寻找下个s[i+1]字符，再从t[j+1]位置找到
// 1. 如果t找完，s子串还有数据，则返回false
// 2. 如果t找完，则返回true
func Method1IsSubSequence(s string, t string) bool {
	// 排查长度过小的
	lens, lent := len(s), len(t)
	if lens > lent {
		return false
	}

	// 子串为空，返回为true
	if s == "" {
		return true
	}

	// 先从s子串选取数据
	var jNextPos int

	// 从子串s[i]选取一个数据，从原始串中idx位置开始找
	for i := 0; i < lens; i++ {
		j := jNextPos
		for {
			// 从t中找s[i]字符时候，找完t也没有发现s[i]，直接返回失败
			if j == lent {
				return false
			}

			// 在j位置找到了s[i]字符
			if s[i] == t[j] {
				jNextPos = j + 1 // 更新下个j的位置
				break
			}

			j++
		}

		// 做了些许优化: 如果在原始串中剩余字符串部分太短，则肯定无法满足找到子串，直接返回false
		if lens-i > lent-j {
			return false
		}

	}

	// 如果子串i已经到头了，表示已经可以再原始串中找到子串
	return true
}

func TestMethod1IsSubSequence(t *testing.T) {

	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"t1", "abc", "ahbgdc", true},
		{"t2", "axc", "ahbgdc", false},
		{"t3", "a", "a", true},
		{"t4", "a", "ba", true},
		{"t5", "a", "aba", true},
		{"t6", "aa", "aba", true},
		{"t7", "baa", "aba", false},
		{"t8", "", "aba", true},
		{"t9", "bb", "ahbgdc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method1IsSubSequence(tt.s, tt.t); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
