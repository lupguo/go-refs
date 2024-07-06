package _0231225

import (
	"fmt"
	"slices"
	"testing"
)

// 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
// 输入：s = "a1b2"
// 输出：["a1b2", "a1B2", "A1b2", "A1B2"]

// 思路: 回溯思路（二叉树）
//  1. base判断
func letterCasePermutation(s string) []string {
	var ret []string
	if s == "" {
		return nil
	}

	// 从第一个字符找起，找到字符位置
	var prevStrs []string
	idx := -1
	for i, c := range s {
		if c >= 'a' && c <= 'z' { // 前缀字符组合，小写变大写
			idx = i
			prevStrs = append(prevStrs, s[:i+1], fmt.Sprintf("%s%c", s[:i], c-32))
			break
		} else if c >= 'A' && c <= 'Z' { // 大写变小写
			idx = i
			prevStrs = append(prevStrs, s[:i+1], fmt.Sprintf("%s%c", s[:i], c+32))
			break
		}
	}
	if idx < 0 { // s中没有找到字母字符
		return []string{s}
	}
	if idx+1 == len(s) {
		return prevStrs
	}

	// 递归从后续子串中拼接结果
	nextStrs := letterCasePermutation(s[idx+1:])
	for _, prevStr := range prevStrs {
		for _, nextStr := range nextStrs {
			ret = append(ret, fmt.Sprintf("%s%s", prevStr, nextStr))
		}
	}

	return ret
}

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []string
	}{
		{"t1", "a", []string{"a", "A"}},
		{"t2", "1", []string{"1"}},
		{"t3", "a1", []string{"a1", "A1"}},
		{"t4", "a1b", []string{"a1b", "a1B", "A1b", "A1B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.str); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}

func TestChar(t *testing.T) {
	t.Logf("%d, %d, %d, %c, %c", 'a', 'A', 'a'-'A', 'A'+32, 'z'-32)
}
