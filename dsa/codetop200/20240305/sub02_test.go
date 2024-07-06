package _0240305

import (
	"testing"
)

func isSubset(a, b []int) bool {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else { // a[i] < b[j]
			return false
		}
	}

	return i == len(a)
}

// a, b :=  []int{1,2,3}, []int{1,1,2,3}
// 判断a是否为b的子集
//  1. 元素个数是否满足
//  2. 双指针i,j分别遍历a,b数组，检测a是否遍历完成
func isChildSet(a, b []int) bool {
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else if a[i] > b[j] { // 当前迭代a的元素比b的元素大，则b的指针应该后移
			j++
		} else { // a[i] < b[j]，因为后续有序，b[j]会越来越大，a[i]肯定无法在b中找到对应元素，直接返回false
			return false
		}
	}

	return false
}

func TestIsChildSet(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want bool
	}{
		{"t1", []int{1}, []int{1}, true},
		{"t2", []int{1}, []int{2}, false},
		{"t3", []int{1, 1}, []int{1}, false},
		{"t4", []int{1, 1}, []int{1, 1}, true},
		{"t5", []int{1, 1}, []int{1, 1, 2}, true},
		{"t6", []int{1, 2}, []int{1, 1, 2}, true},
		{"t7", []int{1, 3}, []int{1, 1, 2, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubset(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
