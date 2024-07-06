package _0230901

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
// 给你两个字符串 word1 和 word2 。
// 请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
func mergeAlternately1(word1 string, word2 string) string {
	w1Len, w2Len := len(word1), len(word2)
	var ss []byte
	if w1Len >= w2Len {
		i := 0
		for ; i < w1Len; i++ {
			if i < w2Len {
				ss = append(ss, word1[i], word2[i])
			} else {
				ss = append(ss, word1[i:]...)
				break
			}
		}
	} else {
		i := 0
		for ; i < w2Len; i++ {
			if i < w1Len {
				ss = append(ss, word1[i], word2[i])
			} else {
				ss = append(ss, word2[i:]...)
				break
			}
		}
	}
	return fmt.Sprintf("%s", ss)

}

// 利用单一游标并行前进
func mergeAlternately(word1 string, word2 string) string {
	w1Len, w2Len := len(word1), len(word2)
	i := 0

	ss := make([]byte, 0, w1Len+w2Len)
	for i < w1Len && i < w2Len {
		ss = append(ss, word1[i], word2[i])
		i++
	}

	if i < w1Len { // word1还有没迭代完的，剩余的补齐
		ss = append(ss, word1[i:]...)
	} else if i < w2Len { // word2还有没迭代完的，剩余的补齐
		ss = append(ss, word2[i:]...)
	}

	return string(ss)
}

func TestSub1768(t *testing.T) {
	type testCase struct {
		name  string
		word1 string
		word2 string
		want  string
	}

	cases := []testCase{
		{"t1", "ab", "cd", "acbd"},
		{"t2", "a", "cd", "acd"},
		{"t3", "ab", "c", "acb"},
	}

	for _, uc := range cases {
		got := mergeAlternately(uc.word1, uc.word2)
		if got != uc.want {
			t.Errorf("user case[%s] got %v, but want %v", uc.name, got, uc.want)
		}
	}
}
