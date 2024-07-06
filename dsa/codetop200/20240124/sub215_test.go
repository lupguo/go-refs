package _0240124

import (
	"container/heap"
	"testing"
)

// 从数组找第K大数字
//
//		思路: 借助swap分区交换函数实现快速交换，知道找到指定值
//	  1. 分区函数: p := partition(nums, lo, hi)
//	  2. 第k大数字: n-k = idx
//	  3. if idx == p return p; if idx > p { p' = partition(nums, p, hi)} else { p' = partition(nums, lo, p)} : 二分查找
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n == 0 || (k <= 0 || k > n) {
		return 0
	}
	wantIdx := n - k
	lo, hi := 0, n-1
	var p int

	// 二分查找
	for lo < hi {
		p = partitionAgain(nums, lo, hi)
		if p == wantIdx {
			break
		} else if p < wantIdx {
			lo = p + 1
		} else {
			hi = p - 1
		}
	}

	return nums[p]
}

// 对nums进行快排分区
func partitionAgain(nums []int, lo int, hi int) int {
	if lo == hi {
		return lo
	}
	// pivot 分区点
	// p := hi
	pivot := nums[hi]

	// 左右游标
	i, j := lo, hi

	for i < j {
		for i < j && nums[i] < pivot {
			i++
		}
		for i < j && nums[j] > pivot {
			j--
		}
		// 游标迭代完成
		if nums[i] == nums[j] {
			break
		}

		// i, j 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	//
	// // 交换pivot 和 i的位置
	// nums[i], nums[p] = nums[p], nums[i]

	// 返回分区点
	return i
}

func TestFindKthLargestAgain(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"t1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"t2", []int{1, 2, 3, 4, 5, 6}, 4, 3},
		{"t3", []int{6, 5, 4, 3, 2, 1}, 1, 6},
		{"t4", []int{3, 2, 4, 2, 1}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)

			if got != tt.output {
				t.Errorf("got %v, but want %v", got, tt.output)
			}
		})
	}
}

func TestPartitionAgain(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		lo     int
		hi     int
		output int
	}{
		{"t1", []int{3, 2, 1, 5, 6, 4}, 0, 5, 3},
		{"t2", []int{1, 2, 3, 4, 5, 6}, 0, 5, 5},
		{"t3", []int{6, 5, 4, 3, 2, 1}, 0, 5, 0},
		{"t4", []int{3, 2, 1}, 0, 2, 0},
		{"t5", []int{3, 2, 2, 1}, 0, 3, 0},
		{"t5", []int{1, 2, 2, 3}, 1, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partitionAgain(tt.nums, tt.lo, tt.hi)

			if got != tt.output {
				t.Errorf("got %v, but want %v", got, tt.output)
			}
		})
	}
}

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素
func findKthLargestV2(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	var mheap MinHeap
	heap.Init(&MinHeap{})
	for _, num := range nums {
		heap.Push(&mheap, num)
		if mheap.Len() > k {
			heap.Pop(&mheap)
		}
	}

	return heap.Pop(&mheap).(int)
}

// 小顶堆方式 -> 构建一个k个元素的小顶堆，然后不断入堆、出堆，最后堆顶元素就是第k大元素
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h MinHeap) PickPop() any {
	return h[h.Len()-1]
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"t1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"t2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"t3", []int{1}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)

			if got != tt.output {
				t.Errorf("got %v, but want %v", got, tt.output)
			}
		})
	}
}
