package _0231012

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75
// 【单词反转】
// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
func reverseWords(s string) string {
	// 1. 清理单词多余的空格（自定义函数，系统函数: 通过正则将多个空格替换成一个; 移除冗余的空格;
	// cleanStr := replaceAllSpaceToSingle(s)
	// ss := strings.Split(cleanStr, " ")

	// // 1. 清理单词多余的空格（自定义函数）
	// ss := cleanRedundancyWords(s)
	// newStr := ""
	//
	// for i := len(ss) - 1; i >= 0; i-- {
	// 	newStr += ss[i] + " "
	// }
	//
	// return strings.Trim(newStr, " ")

	return reverseWordsByIJ(s)
}

// 解法3: 双游标，通过s[m:n] 从后向前提取字符串
func reverseWordsByIJ(s string) string {
	var words []string
	// 外循环，从后前迭代游标
	for i := len(s) - 1; i >= 0; i-- {
		// 原始是空字符，继续向前
		if s[i] == ' ' {
			continue
		}

		// 找到了一个非0字符，开启内循环找单词
		j := i
		for ; j >= 0; j-- {
			// 字符迭代到空的情况，则切片 s[j+1:i+1]为要找的word
			if s[j] == ' ' {
				words = append(words, s[j+1:i+1])
				break
			}
			// 字符迭代到第一个字符了，则切片s[j:i+1]为要找的word
			if j == 0 {
				words = append(words, s[j:i+1])
			}
		}

		// 游标向前移动
		i = j
	}

	return strings.Join(words, " ")
}

// 解法1
func replaceAllSpaceToSingle(s string) string {
	rgx := regexp.MustCompile(`\s+`)
	return rgx.ReplaceAllString(strings.Trim(s, " "), " ")
}

// 解法2: 清理单词多余的空格
func cleanRedundancyWords(s string) []string {
	// 遍历字符串s的每个字符，提取每个元素，放入新的字符切片中
	var word string
	var words []string
	for i, c := range s {
		// 空字符，将非空的word装载起来，重置word
		if c == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			continue
		}

		// 非空字符先
		word = fmt.Sprintf("%s%c", word, c)

		// 字符是否为最后一个字符，且word不为空
		if i == len(s)-1 && word != "" {
			words = append(words, word)
		}
	}

	return words
}

func TestCleanRedundancyWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"t1", "hello world", []string{"hello", "world"}},
		{"t2", "hello  world", []string{"hello", "world"}},
		{"t3", " hello world", []string{"hello", "world"}},
		{"t4", "hello  world ", []string{"hello", "world"}},
		{"t5", "hey, hello  world ", []string{"hey,", "hello", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanRedundancyWords(tt.s); reflect.DeepEqual(got, tt.want) != true {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"t1", "hey man", "man hey"},
		{"t2", "hey  man", "man hey"},
		{"t3", " hey  man", "man hey"},
		{"t4", " hey  man ", "man hey"},
		{"t5", "hi, hey  man ", "man hey hi,"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("got `%v`, but want `%v`", got, tt.want)
			}
		})
	}

}

func TestReplaceAllSpaceToSingle(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"t1", "hello world", []string{"hello", "world"}},
		{"t2", " hello world", []string{"hello", "world"}},
		{"t3", " hello   world", []string{"hello", "world"}},
		{"t4", " hello   world  ", []string{"hello", "world"}},
		{"t5", "hi,   hello   world  ", []string{"hi,", "hello", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceAllSpaceToSingle(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got `%#v`, but want `%#v`", got, tt.want)
			}
		})
	}
}
