package _0240407

import (
	"testing"
)

// 查找a, b两个有序数组的中位数
// 思路:
//  1. 获取a,b 数组长度，lenA,LenB，因为有序，所以中位数应该是在 (lenA + lenB) / 2 位置
//  2. i, j 分别从a, b 数组迭代, a[i], b[j] 同步向前迭代，需要比较a[i], b[j]大小
//     直到遇到 i+j = (lenA + lenB) / 2  停止返回
func findMidNum(a, b []int) int {
	lenA, lenB := len(a), len(b)
	midIdx := (lenA + lenB) / 2
	var i, j int

	// 迭代到中位数索引位置
	var midNum int
	for k := 0; k <= midIdx; k++ {
		// i, j 比较大小，逐步向前遍历
		if i < lenA && (j == lenB || a[i] < b[j]) {
			midNum = a[i]
			i++
		} else {
			midNum = b[j]
			j++
		}
	}

	return midNum
}

func TestFindMidNum(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want int
	}{
		{"t1", []int{1}, []int{2}, 2},
		{"t2", []int{1, 2}, []int{3}, 2},
		{"t3", []int{1, 2}, []int{3, 4}, 3},
		{"t4", []int{1, 2}, []int{3, 4}, 3},
		{"t5", []int{3, 4}, []int{1, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMidNum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
