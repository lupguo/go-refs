package _0231116

import (
	"testing"
)

// 找到最高海拔
// 有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。
// 示例 1：
//
// 输入：gain = [-5,1,5,0,-7]
// 输出：1
// 解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
// 示例 2：
//
// 输入：gain = [-4,-3,-2,-1,4,3,2]
// 输出：0
// 解释：海拔高度依次为 [0,-4,-7,-9,-10,-6,-3,-1] 。最高海拔为 0 。
func largestAltitude(gain []int) int {

	return 0
}

// 方法2: 前缀和，不用初始数组
// haiba[0] = 0 (i <1)
// haiba[i] = gain[i-1] + haiba[i-1]，取最大haiba[i] (i >=1)
func LargestAltitudeMethod02(gain []int) int {
	maxHaiba := 0
	curHaiba := 0

	// i从(1,n),
	for i := 1; i < len(gain)+1; i++ {
		curHaiba = gain[i-1] + curHaiba
		maxHaiba = max(curHaiba, maxHaiba)
	}

	return maxHaiba
}

// 解题思路
// 1. 净海拔高度gain差 => 推导出 [0,n+1] 海拔高度数组
// 2. 从海拔高度数组中选取最高的海拔返回
// 从0海拔出发，gain[i-1] = haiba[i] - haiba[i-1] => haiba[1] = gain[0] + haiba[0]
func LargestAltitudeMethod01(gain []int) int {
	n := len(gain)
	haiba := make([]int, n+1)
	var maxHaiba int
	for i := 1; i <= n; i++ {
		haiba[i] = haiba[i-1] + gain[i-1]
		maxHaiba = max(haiba[i], maxHaiba)
	}

	return maxHaiba
}

func TestLargestAltitudeMethod01(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{"t1", []int{-5, 1, 5, 0, -7}, 1},
		{"t2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestAltitudeMethod02(tt.gain); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
