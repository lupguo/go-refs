package _0231121

import (
	"slices"
	"testing"
)

// 如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：
//
// 操作 1：交换任意两个 现有 字符。
// 例如，abcde -> aecdb
// 操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
// 例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
// 你可以根据需要对任意一个字符串多次使用这两种操作。
//
// 给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：word1 = "abc", word2 = "bca"
// 输出：true
// 解释：2 次操作从 word1 获得 word2 。
// 执行操作 1："abc" -> "acb"
// 执行操作 1："acb" -> "bca"
//
// 示例 2：
//
// 输入：word1 = "a", word2 = "aa"
// 输出：false
// 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
//
// 示例 3：
//
// 输入：word1 = "cabbba", word2 = "abbccc"
// 输出：true
// 解释：3 次操作从 word1 获得 word2 。
// 执行操作 1："cabbba" -> "caabbb"
// 执行操作 2："caabbb" -> "baaccc"
// 执行操作 2："baaccc" -> "abbccc"
// 示例 4：
//
// 输入：word1 = "cabbba", word2 = "aabbss"
// 输出：false
// 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
func closeStrings(word1 string, word2 string) bool {

	return true
}

// 思路: 实际是统计两个字符串中字符计数是否一致
// 1. word1, word2 -> m1, m2 map[byte]int
// 2. m1,m2 -> arr1, arr2
// 3. sort(arr1), sort(arr2)
// 4. compare(arr1', arr2')
// --
// 注意: 需要检测元素Key是否在内哈
func CloseStringsMethod1(word1 string, word2 string) bool {
	// 1. 转map计数
	m1, m2 := make(map[rune]int), make(map[rune]int)
	for _, c := range word1 {
		m1[c]++
	}
	for _, c := range word2 {
		m2[c]++
	}

	// 2. 转数组
	var arr1, arr2 []int
	for _, cnt := range m1 {
		arr1 = append(arr1, cnt)
	}
	slices.Sort(arr1)

	for c, cnt := range m2 {
		arr2 = append(arr2, cnt)

		// 检测m1的元素是否都在m2内
		if _, ok := m1[c]; !ok {
			return false
		}
	}
	slices.Sort(arr2)

	// 3. sort后比较
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestCloseStringsMethod1(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		{"t1", "abc", "cba", true},
		{"t2", "cabbba", "aabbss", false},
		{"t3", "a", "aa", false},
		{"t4", "cabbba", "abbccc", true},
		{"t5", "uau", "ssx", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloseStringsMethod1(tt.word1, tt.word2); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
