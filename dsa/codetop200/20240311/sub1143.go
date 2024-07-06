package _0240311

// longestCommonSubsequence
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
func longestCommonSubsequenceV1(text1 string, text2 string) int {
	return 0
}

// 思路: 从底部向上，想得到dp[i][j]元素，假定dp[i][j]以前的元素已经得到，可以由它们推导而来
//   - dp[i][j] = 1+dp[i-1][j-1] // s1[i] == s1[j]情况
//   - dp[i][j] = max(dp[i][j-1], dp[i-1][j]) //
func longestCommonSubsequence(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	// 初始化dp迭代参数
	lcs := make([][]int, m)
	for i := 0; i < m; i++ {
		lcs[i] = make([]int, n)
		for j := 0; j < j; j++ {
			// 相等
			if s1[i] == s2[j] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	// 返回迭代结果
	return lcs[m][n]
}

// 返回两个字符串的最长公共子序列
//
//	abc bac -> ac|bc -> 2
//
// 思路: 递归思路，至顶向下
//
//	fn(s1,i,s2,j) =
//	    * fn(s1,i+1, s2, j+1) // 若 s1[i] == s2[j]
//	    * max(fn(s1,i+1, s2, j), fn(s1,i, s2,j+1)) // 否则
func longestCommonSubsequenceV3(s1, s2 string) int {
	// 初始化暂存信息
	mem := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		mem[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			mem[i][j] = -1
		}
	}

	// DP递归
	return dpV3(s1, 0, s2, 0, mem)
}

func dpV3(s1 string, i int, s2 string, j int, mem [][]int) int {

	// base case
	if s1 == "" || s2 == "" {
		return 0
	} else if i >= len(s1) || j >= len(s2) {
		return 0
	}

	// i, j是否有缓存
	if mem[i][j] != -1 {
		return mem[i][j]
	}

	// mem无缓存，从顶向下递归
	if s1[i] == s2[j] {
		// 保存到缓存
		mem[i][j] = 1 + dpV3(s1, i+1, s2, j+1, mem)
	} else {
		mem[i][j] = max(
			dpV3(s1, i, s2, j+1, mem),
			dpV3(s1, i+1, s2, j, mem),
		)
	}

	return mem[i][j]
}
