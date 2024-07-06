package _0240312

// https://leetcode.cn/problems/word-break/description/
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

// 递归思路: dp(s, i, words)，从i位置迭代s，探测s[i:j]是否再words中，若在
//   - 若在，问题变为: dp(s, j, words)
func wordBreak(s string, wordDict []string) bool {
	// 将wordDict转成map，方便后续比较
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	// 增加缓存优化
	mem := make([]int, len(s))

	return dp(s, 0, dict, mem)
}

// 检测字符串s从下标i位置是否存在s[i:j]在dict内，并且剩余的
func dp(s string, i int, dict map[string]bool, mem []int) bool {
	if i >= len(s) {
		return true
	}
	if mem[i] > 0 {
		return true
	}

	// 递归从s中查找前缀str是否在
	for j := i; j <= len(s); j++ {
		str := s[i:j]
		if dict[str] { // s的前缀str字符串可以在dict内能找到
			found := dp(s, j, dict, mem)
			if found {
				mem[i] = 1
				return true
			}
		}
	}

	return false
}
