package _0240311

// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和
// 思路:
//
//	mds[i][j] = mds(s1, i+1, s2, j+1) // 两者相等
//	 mds[i][j] = min('s1[i]'+mds(s1,i+1,s2,j), 's2[j]'+mds(s1,i,s2,j+1)) // 要么删除s1[i]，要么删除s2[j]再继续看两者要删除的公共字符ascii
func minimumDeleteSum(s1 string, s2 string) int {

	// 初始化mem(-1的二维数组）
	m, n := len(s1), len(s2)
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}

	// dp动态规划
	return mds(s1, 0, s2, 0, mem)
}

// 获取s1[i:], s2[j:]需要删除ascii和
func mds(s1 string, i int, s2 string, j int, memo [][]int) int {
	// base case - i,j 超出，则另外一个字符串需要全部删除
	if i >= len(s1) {
		return sumAscii(s2[j:])
	} else if j >= len(s2) {
		return sumAscii(s1[i:])
	}

	// 是否已经mem有值了，缓存加速
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// 动态递归后续最优解
	if s1[i] == s2[j] {
		memo[i][j] = mds(s1, i+1, s2, j+1, memo) // s1[i],s2[j]相等，继续向后移动
	} else {
		memo[i][j] = min(
			int(s1[i])+mds(s1, i+1, s2, j, memo), // 删除i, s1 i->i+1
			int(s2[j])+mds(s1, i, s2, j+1, memo), // 删除j, s2 j->j+1
		)
	}

	return memo[i][j]
}

func sumAscii(s string) int {
	var n int
	for _, r := range s {
		n += int(r)
	}
	return n
}
