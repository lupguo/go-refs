package _0240311

// https://leetcode.cn/problems/climbing-stairs/description/
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//	思路1.
//	    自顶向下递归：爬到n方法种数 claw(n)= claw(n-1)+claw(n-2), base n=1 ret=1,n=2 ret=2
func climbStairs(n int) int {
	mem := make([]int, n+1)
	return climb(n, mem)
}

func climb(n int, mem []int) int {
	if n <= 2 {
		return n
	}

	// 缓存
	if mem[n] != 0 {
		return mem[n]
	}

	mem[n] = climb(n-1, mem) + climb(n-2, mem)
	return mem[n]
}

// 自底向上
func climbStairsV2(n int) int {
	// base case
	if n <= 2 {
		return n
	}

	// 自低向上迭代
	sum := make([]int, n+1)
	sum[0], sum[1] = 1, 1
	for i := 2; i <= n; i++ {
		sum[i] = sum[i-1] + sum[i-2]
	}

	return sum[n]
}

// 自底向上 v3
func climbStairsV3(n int) int {
	// base case
	if n <= 2 {
		return n
	}

	// 自低向上迭代
	p, q, r := 1, 1, 0
	for i := 2; i <= n; i++ {
		r = p + q
		p, q = q, r
	}

	return r
}
