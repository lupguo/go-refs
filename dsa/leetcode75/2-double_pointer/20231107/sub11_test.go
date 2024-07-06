package _0231107

import (
	"reflect"
	"testing"
)

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
func maxArea(height []int) int {

	return Method1MaxArea(height)
}

// area(i,j)=min(h[i],h[j])*(j-i)
// 因为j-i可以选择，i,j 分别放在两端指定位置，打算向中间夹逼选择最大值
// 1. 这样当移动i或者j时候, j-i一定会变小
// 2. 盛水量是由短板决定，短边限定了增大可能性，所以优先移动短板
// 时间复杂度 O(N) : 每个板子都移动了一次
// 空间复杂度 O(1）: d
func Method2MaxArea(height []int) int {
	n := len(height)
	i, j := 0, n-1
	var maxArea int
	// 左指针右移，右指针左移
	for i < j {
		curArea := min(height[i], height[j]) * (j - i)
		if curArea > maxArea {
			maxArea = curArea
		}

		// 思考下该移动哪个板子 -> 按盛水面积，下次迭代(j-i)一定会缩小，那么优先移动短板才是正确选择
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

// 暴力算法 - 容易超出时间限制
func Method1MaxArea(height []int) int {
	n := len(height)
	// j := n - 1
	maxAreaSize := 0
	// i 从左往右移动
	for i := 0; i < n-1; i++ {

		// // j 从右向左移动
		// for j = n - 1; i < j && j > 0; j-- {
		// 	ijArea := min(height[i], height[j]) * (j - i)
		// 	if ijArea > maxAreaSize {
		// 		maxAreaSize = ijArea
		// 	}
		// }

		// 或者选择从左i+1向右推进，依次计算area(i,j)中面积最大值
		for j := i + 1; j < n; j++ {
			ijArea := min(height[i], height[j]) * (j - i)
			if ijArea > maxAreaSize {
				maxAreaSize = ijArea
			}
		}
	}

	return maxAreaSize
}

func TestMethod1MaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"t1", []int{1, 1}, 1},
		{"t2", []int{1, 2}, 1},
		{"t3", []int{2, 1}, 1},
		{"t4", []int{1, 2, 1}, 2},
		{"t5", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method2MaxArea(tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}

}
