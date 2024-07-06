package _0231011

import (
	"testing"
)

// https://leetcode.cn/problems/can-place-flowers/description/?envType=study-plan-v2&envId=leetcode-75
//
// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
func canPlaceFlowers(flowerbed []int, n int) bool {
	// 针对左边界条件，直接补[1,0]，针对右边界，补[0,1]
	zeroCnt := 0
	maxAllowCnt := 0
	for i, v := range flowerbed {
		// 左边界为0，默认zero的count计数应该增加1
		if i == 0 && v == 0 {
			zeroCnt++
		}

		// 两个种花为1数值的中间位置
		// 检测i位置v的值，如果为0，则计数0加1，同时继续往后检测
		// 如果非0，则将统计两个1之间的0可种花数量，同时直接将计数0重置，并跳过继续下一轮检测
		if v == 0 {
			zeroCnt++
		} else {
			maxAllowCnt += (zeroCnt - 1) / 2
			zeroCnt = 0
			continue
		}

		// 右边界为0，默认zero的count计数增加1，同时补充上允许种花的值
		if i == len(flowerbed)-1 && v == 0 {
			zeroCnt++
			maxAllowCnt += (zeroCnt - 1) / 2
		}
	}

	return maxAllowCnt >= n
}

func Test_CanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{"t1", []int{1, 0, 0, 0, 1}, 1, true},
		{"t2", []int{1, 0, 0, 0, 1}, 2, false},
		{"t3", []int{0, 0}, 1, true},
		{"t4", []int{0, 0}, 2, false},
		{"t5", []int{0, 0, 0}, 2, true},
		{"t6", []int{0, 1, 0}, 1, false},
		{"t7", []int{0, 1}, 1, false},
		{"t8", []int{0}, 1, true},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if got := canPlaceFlowers(c.flowerbed, c.n); got != c.want {
				t.Errorf("case[%v] got %v, but want %v", c.name, got, c.want)
			}
		})
	}

}
