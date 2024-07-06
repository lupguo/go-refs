package sortx

import (
	"slices"
	"testing"
)

// 快排思路:
//  1. 先排好指定位置q，再递归的排序nums[p, q],  nums[q],  nums[q+1, r]部分
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 找到分区点，并将q位置排好序
	q := partition(nums)

	// 递归分区间进行子序列的排序
	quickSort(nums[0:q])
	quickSort(nums[q+1:])

	return nums
}

// 快排的分区函数:
//
//	思路: 设定n-1为pivot分区点，基于pivot值将列表数据分层两部分：
//	    - 划分成已处理[0:i-1] 和 未处理 [j:n-2]两个区间
//	    - 未处理的如果满足 nums[j] <= pivot，则i++, j++，
//	        若 num[j] > pivot，对swap(num[i], num[j])交换，使之满足 num[j] < pivot成立
func partition(nums []int) int {
	// 只有一个元素 或者 空元素则直接返回切片长度
	r := len(nums)
	if r <= 1 {
		return r
	}

	// 默认最后一个元素为分区点
	pivot := nums[r-1]
	var i int

	// j 就是一个不断去找小的探针
	for j := 0; j < r-1; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
	}
	nums[i], nums[r-1] = nums[r-1], nums[i]

	return i
}

func TestQuickSort(t *testing.T) {
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
			if got := quickSort(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})

	}
}
