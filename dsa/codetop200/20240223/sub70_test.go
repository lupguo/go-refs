package _0240223

import (
	"testing"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 从上到下递归思考问题，将原问题拆解层子问题
func climbStairs2(n int) int {
	save := make(map[int]int)
	var dp func(n int) int
	dp = func(n int) int {
		// 检测释放已有爬楼记录
		if v, ok := save[n]; ok {
			return v
		}

		// base
		if n <= 2 {
			return n
		}

		// 状态转移方程, 保存爬楼记录
		save[n] = dp(n-1) + dp(n-2)
		return save[n]
	}

	return dp(n)
}

// https://leetcode.cn/problems/climbing-stairs/description/
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs1(n int) int {
	// 爬楼方法
	var method int
	save := make(map[int]int)

	// 定义dp函数: 爬n楼到方法次数
	var dp func(n int, method *int)
	dp = func(n int, method *int) {
		// 已有爬的存储记录
		if v, ok := save[n]; ok {
			*method += v
			return
		}

		// base - 找到了一种爬楼方法
		if n == 0 {
			*method++
			return
		}

		// 每次选择可能
		for _, step := range []int{1, 2} {
			// 选择支持
			if step > n {
				continue
			}

			// 子问题递归
			dp(n-step, method)
		}

		// 保存n部
		save[n] = *method
	}
	dp(n, &method)

	return method
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"t1", 1, 1},
		{"t2", 2, 2},
		{"t3", 3, 3},
		{"t4", 4, 5},
		{"t5", 5, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
