package _0240407

import (
	"testing"
)

// 查找a, b两个有序数组的中位数
// 思路:
//  1. 获取a,b 数组长度，lenA,LenB，因为有序，所以中位数应该是在 (lenA + lenB) / 2 位置
//  2. i, j 分别从a, b 数组迭代, a[i], b[j] 同步向前迭代，需要比较a[i], b[j]大小
//     直到遇到 i+j = (lenA + lenB) / 2  停止返回
//
// 收获：
//  1. 迭代a,b数组到中位数，可以借助第三个变量k迭代
//  2. 奇偶数，可以考虑使用暂存前一个midNum方式，最后再通过奇偶判断
func findMedianSortedArrays(a, b []int) float64 {
	lenA, lenB := len(a), len(b)

	// 中间游标位置
	midIdx := (lenA + lenB) / 2

	// 两个游标i, j
	var i, j, tmp int
	var midNum int
	for k := 0; k <= midIdx; k++ {
		tmp = midNum
		// i,j 迭代
		if i < lenA && (j == lenB || a[i] < b[j]) {
			midNum = a[i]
			i++
		} else {
			midNum = b[j]
			j++
		}
	}

	if (lenA+lenB)%2 == 0 { // 偶数
		return float64(midNum+tmp) / 2
	}

	return float64(midNum)
}

func TestFindMidNum2(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want float64
	}{
		{"t1", []int{1}, []int{2}, 1.5},
		{"t2", []int{1, 2}, []int{3}, 2},
		{"t3", []int{1, 2}, []int{3, 4}, 2.5},
		{"t4", []int{1, 2}, []int{3, 4}, 2.5},
		{"t5", []int{3, 4}, []int{1, 2}, 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
