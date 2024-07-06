package _0240106

import (
	"testing"
)

// 结雨水
// 思路:
//
//	得到储水公式: storageWater[i] = min(leftMaxHeight[i], rightMaxHeight[i]) - curHeight[i]
func trap(height []int) int {
	// 1. 先分别计算下 leftMaxHeight、rightMaxHeight[i]的值, 分别存储到第i位置(不包含)的最大高度
	n := len(height)
	leftMaxHeight, rightMaxHeight := make(map[int]int, n), make(map[int]int, n)

	// 左侧最大高度初始化
	leftMaxHeight[0] = 0
	for i := 1; i < n; i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i-1])
	}
	// 右侧最大高度初始化
	rightMaxHeight[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i+1])
	}

	// 2. 循环计算储水量
	var storageWaterTotal int
	for i := 0; i < n; i++ {
		if storageWater := min(leftMaxHeight[i], rightMaxHeight[i]) - height[i]; storageWater > 0 {
			storageWaterTotal += storageWater
		}
	}

	return storageWaterTotal
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"t1", []int{0, 1}, 0},
		{"t2", []int{1, 1}, 0},
		{"t3", []int{1, 2, 1}, 0},
		{"t4", []int{1, 0, 1}, 1},
		{"t5", []int{1, 0, 2, 0, 1}, 2},
		{"t6", []int{0, 0, 2, 0, 1}, 1},
		{"t7", []int{1, 0, 2, 0, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
