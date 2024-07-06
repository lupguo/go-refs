package week12

import (
	"container/heap"
)

// https://leetcode.cn/problems/minimum-operations-to-halve-array-sum/
// 随机选择数组中数字，折半到数组和一半至少操作
// 输入：nums = [5,19,8,1]
// 输出：3
// 思路:
//  1. 求一半和 target <= halfSum
//  2. 大顶堆
//  3. 依次选择堆顶元素，将折半后内容继续加入堆中堆化处理
func halveArray(nums []int) int {
	// 1. 求和
	n := len(nums)
	var sum int
	fnums := make([]float64, n)
	for i, num := range nums {
		sum += num
		fnums[i] = float64(num)
	}
	halfSum := float64(sum) / 2

	// 2. 大顶堆
	hp := NewBigHeap(fnums)
	heap.Init(hp)

	// 3. 迭代选择减半或者是选择最大值
	var dedTimes int    // 扣减次数
	var dedSum float64  // 扣减和
	var halfVal float64 // 取半值

	for dedSum < halfSum {
		// 每次选择最大值，统计扣减总和
		halfVal = hp.Pop().(float64) / 2
		dedSum += halfVal

		// 减掉半值相加
		dedTimes++
		hp.Push(halfVal)
	}

	return dedTimes
}

type BigHeap []float64

func NewBigHeap(nums []float64) *BigHeap {
	return (*BigHeap)(&nums)
}

func (h *BigHeap) Len() int {
	return len(*h)
}

func (h *BigHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *BigHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *BigHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *BigHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
