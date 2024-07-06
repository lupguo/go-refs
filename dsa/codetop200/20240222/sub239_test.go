package _0240222

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/sliding-window-maximum/

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
// 思路:
//  1. left,right, window k大小窗口滑动
//  2. 窗口移动，获取其中最大值（可以基于单调队列做O(1) 取最大值优化）
func maxSlidingWindow(nums []int, k int) []int {

	// 边界
	length := len(nums)
	if k > length {
		return nil
	}

	// 扩展到k窗口
	var windows MonoQueue
	for i := 0; i < k; i++ {
		windows.Push(nums[i])
	}

	// 窗口更新&获取k窗口最大值
	var ret []int
	left, right := 0, k-1
	for right < length {
		// 窗口处理
		ret = append(ret, windows.Max()) // 可以使用单调队列优化

		// 窗口更新
		right++
		if right < length {
			windows.Push(nums[right])
			windows.Pop(nums[left])
		}
		left++
	}

	return ret
}

// 单调递增队列
type MonoQueue struct {
	data []int
	max  int
}

// 单调队列元素入队，保持递减
func (q *MonoQueue) Push(num int) {
	// 入队元素比队当前队尾元素大，则需要更新单调队列保持 递减属性
	length := len(q.data)
	for length > 0 && q.data[length-1] < num {
		q.data = q.data[:length-1]
		length = len(q.data)
	}
	// 元素入队
	q.data = append(q.data, num)

	// // 更新max值
	// if q.data[len(q.data)-1] != q.max {
	// 	q.max = q.data[0]
	// }
}

// 获取单调递增Max元素，队首为最大元素
func (q *MonoQueue) Max() int {
	return q.data[0]
}

// 单调队列出队，如果出队元素和队首相同，则更新单调队列
func (q *MonoQueue) Pop(num int) {
	if len(q.data) > 0 && num == q.data[0] {
		q.data = q.data[1:]
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "t1", nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, want: []int{3, 3, 5, 5, 6, 7}},
		{name: "t2", nums: []int{1}, k: 1, want: []int{1}},
		{name: "t3", nums: []int{9, 11}, k: 2, want: []int{11}},
		{name: "t4", nums: []int{}, k: 3, want: nil},
		{name: "t5", nums: []int{1, -1}, k: 1, want: []int{1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %v, but want %v", result, tt.want)
			}
		})
	}
}
