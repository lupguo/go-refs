package _0240314

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// [0, 1, 0, 2, 1, 0
// [i, j]
//
//	[i, j, j]
//
// 任意一个idx下标i位置可以装水 water[i] = min(maxHeight(height[0:i]), maxHeight(height[i+1:]) - height[i]
// 考虑: 边界是魔鬼（可以画图举例），弄清楚
//   - leftMaxHeight[i]、rightMaxHeight[i]含义
//   - 记录公式
func trap(height []int) int {
	n := len(height)

	// 左侧最高
	leftMaxHigh := make([]int, n+1)
	for i := 0; i < n; i++ {
		// leftMaxHigh[i] 表示在i位置（不包含i) 的左侧最大高度
		leftMaxHigh[i+1] = max(height[i], leftMaxHigh[i])
	}
	// 右侧最高
	rightMaxHigh := make([]int, n+1)
	for i := n - 1; i >= 1; i-- {
		// rightMaxHigh[i] 表示在i位置右侧（不包含i) 的左侧最大高度
		rightMaxHigh[i-1] = max(height[i], rightMaxHigh[i])
	}

	// 储水量
	var total int
	for i := 0; i < n; i++ {
		// 计算每个idx下标位置最大可储水量
		water := min(leftMaxHigh[i], rightMaxHigh[i]) - height[i]
		if water > 0 {
			total += water
		}
	}

	return total
}
