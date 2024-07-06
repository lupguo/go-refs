package _0231226

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 从原始串，找子串的组合问题 -> 滑动窗口，返回满足条件的 left 左边框数组集合
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
func findAnagrams(s string, p string) []int {
	// 异常case
	if len(p) > len(s) {
		return nil
	}

	// 目标窗口
	needWin := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		needWin[p[i]]++
	}

	var ret []int
	curWin := make(map[byte]int) // 当前窗口
	left, right := 0, 0          // 窗口左右边框

	// 窗口滑动
	for right < len(s) {
		// 右滑窗口
		curWin[s[right]]++
		right++

		// 目前窗口过小，继续右滑
		if right-left < len(p) {
			continue
		}

		// 窗口合适，窗口比较看是否符合预期
		if isMeetWindow(curWin, needWin) { // 找到合适子串，将左边框位置加入返回值
			ret = append(ret, left)
		}

		// 左窗口右移，缩小窗口大小，更新窗口内容
		curWin[s[left]]--
		left++
	}

	return ret
}
