package _02404

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 滑动窗口思路, i,j
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	set := make(map[byte]bool)
	maxLen, i, j := 0, 0, 0

	for i < n && j < n {
		if !set[s[j]] {
			set[s[j]] = true
			j++
			maxLen = max(maxLen, j-i)
		} else {
			delete(set, s[i])
			i++
		}
	}

	return maxLen
}

func winCheckOk(b byte, window []byte) bool {
	for _, c := range window {
		if b == c {
			return true
		}
	}
	return false
}
