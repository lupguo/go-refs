package sortx

import (
	"slices"
	"testing"
)

// mergeSort 实现一个归并排序算法
//
//	思路: 分治+合并
//	    1. 对数组分区间q=r-p/2 , left := mergeSort([p,q], p, q), right := mergeSort([q+1,r], q+1, r)
//	    2. 合并两个有序数组
func mergeSort(nums []int) []int {
	// 基准case
	n := len(nums)
	if n <= 1 {
		return nums
	}

	// 分治
	q := n / 2
	left := mergeSort(nums[0:q])
	right := mergeSort(nums[q:n])

	// 合并两个有序数组
	return merge(left, right)
}

// 合并两个有序数组
func merge(left []int, right []int) []int {
	llen, rlen := len(left), len(right)

	// 两个游标i,j分别在left, right数组上游动
	var ret []int
	var i, j int
	for i < llen && j < rlen {
		// 将i,j元素中更小的放入ret切片中
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}

	// 检测i,j哪个已经迭代完成了，将剩余的最佳到ret中
	if i == llen {
		ret = append(ret, right[j:]...)
	} else {
		ret = append(ret, left[i:]...)
	}

	return ret
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"t1", []int{1}, []int{1}},
		{"t2", []int{2, 1}, []int{1, 2}},
		{"t3", []int{2, 1, 3}, []int{1, 2, 3}},
		{"t4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"t5", []int{3, 2, 1, 4}, []int{1, 2, 3, 4}},
		{"t6", []int{3, 2, 1, 3}, []int{1, 2, 3, 3}},
		{"t7", []int{3, 1, 1, 3}, []int{1, 1, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})

	}
}
