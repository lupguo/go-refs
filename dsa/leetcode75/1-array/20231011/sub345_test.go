package _0231011

import (
	"testing"
)

// https://leetcode.cn/problems/reverse-vowels-of-a-string/?envType=study-plan-v2&envId=leetcode-75
// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
func reverseVowels(s string) string {
	// 两边夹逼，先找到左侧元音，再找到右侧元音

	i, j := 0, len(s)-1
	revBytes := []byte(s)
	for i < j {
		// 开始迭代左侧元素，如果找到元音则停止，记录下标到left中，跳出左侧循环
		for !isYuanYin(revBytes[i]) && i < j {
			i++
		}

		// 开始右侧元素迭代，如果找到元音则停止，记录右侧下标到right中，跳出右侧循环
		for !isYuanYin(revBytes[j]) && j > 0 {
			j--
		}

		// 如果左侧下标小于右侧下标，则进行交换
		if i < j {
			revBytes[i], revBytes[j] = revBytes[j], revBytes[i]
			i++
			j--
		}
	}

	return string(revBytes)
}

var yuanYinMap = map[byte]bool{
	'a': true,
	'A': true,
	'e': true,
	'E': true,
	'i': true,
	'I': true,
	'o': true,
	'O': true,
	'u': true,
	'U': true,
}

func isYuanYin(c byte) bool {
	return yuanYinMap[c]
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"t1", "he", "he"},
		{"t2", "heo", "hoe"},
		{"t3", "hello", "holle"},
		{"t4", "helloe", "helloe"},
		{"t5", "helloeo", "holleoe"},
		{"t6", "aA", "Aa"},
		{"t7", "", ""},
		{"t8", ".,", ".,"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("got %s, but want %s", got, tt.want)
			}
		})
	}
}
