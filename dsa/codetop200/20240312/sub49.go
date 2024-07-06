package _0240312

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
//
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
//	输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//	输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 思路:
//  1. 分组，统计每个str的字母计数统计情况，一致的字符串则表述可以进行异位
//     * eat: map[eat].charCount == map[ate].charCount => append(eat, ate)
//     *

// 分组异位字符串
// 思路，如何快速判断两个字符串异位 => encode编码成str, eat=>e1a1t1 tae=>t1a1e1
func groupAnagrams(strs []string) [][]string {

	// 尝试将strs中的字符串转成同一异位编码的map元素中
	m := make(map[string][]string)
	for _, str := range strs {
		enc := encodeStr(str)
		m[enc] = append(m[enc], str)
	}

	// 迭代enc map返回结果
	var ans [][]string
	for _, ss := range m {
		ans = append(ans, ss)
	}

	return ans
}

// 利用每个字符的出现次数进行编码
func encodeStr(s string) string {
	count := make([]byte, 26) // 使用切片替代map，s[i]-'a'统计位置
	for i := 0; i < len(s); i++ {
		delta := s[i] - 'a' // delta作为key
		count[delta]++      // 字符计数
	}

	return string(count) // 二进制串
}

// 返回str的有序的字符统计串: tea => a1e1t1
func encodeStr1(str string) string {
	cs := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		cs[str[i]]++
	}

	// 转成slice字符串
	var ss []string
	for s, cnt := range cs {
		ss = append(ss, fmt.Sprintf("%c%d", s, cnt))
	}

	// 对ss排序
	sort.Strings(ss)

	// 返回结果
	return strings.Join(ss, "")
}

// CharMap 字符计数map
type CharMap map[byte]int

// Equal 两个charMap是否互为异位
func (cm CharMap) Equal(m CharMap) bool {
	for char, cnt := range cm {
		if m[char] != cnt {
			return false
		}
	}
	return true
}

// StrToCharMap 字符串转charMap
func StrToCharMap(str string) CharMap {
	cm := make(CharMap)
	for i := 0; i < len(str); i++ {
		cm[str[i]]++
	}
	return cm
}

// 分组异位字符串
func groupAnagramsV1(strs []string) [][]string {
	var ret [][]string

	// 字符串切片数组，转成charMap分组
	strCharMap := make(map[string]CharMap)
	for _, str := range strs {
		strCharMap[str] = StrToCharMap(str)
	}

	// 遍历strs数组，转存map比较
	used := make(map[string]bool)
	for _, str1 := range strs {
		// str1已被使用
		if used[str1] {
			continue
		}

		ans := []string{str1}
		used[str1] = true
		for _, str2 := range strs {
			// 通用的str跳过，长度不一样的跳过
			if used[str2] || len(str1) != len(str2) {
				continue
			}

			// str1 和 str2互为异位
			if strCharMap[str1].Equal(strCharMap[str2]) {
				ans = append(ans, str2)
				used[str2] = true
			}
		}
		ret = append(ret, ans)
	}

	return ret
}
