package _0231127

import (
	"testing"

	. "dsa/data-struct"
)

// 思路: 前半段反转链表(leftLink)，后半段(rightLink)，两个链表同时迭代，计算元素最大和
//  1. 如何确定到了链表中间位置？ 快慢指针（1步,2步)
func PairSumMethod02(head *ListNode) int {
	var leftLink, rightLink *ListNode
	var prev *ListNode
	slow, fast := head, head.Next

	// 反转前半部分（快慢指针）
	for fast.Next != nil {
		fast = fast.Next.Next

		// 反转
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}

	// 后半段链表
	rightLink = slow.Next
	// 前半段反转链表
	slow.Next = prev
	leftLink = slow

	// 遍历两个链表
	var maxSum int
	for leftLink != nil {
		curSum := leftLink.Val + rightLink.Val
		if curSum > maxSum {
			maxSum = curSum
		}
		leftLink = leftLink.Next
		rightLink = rightLink.Next
	}

	return maxSum
}

// 思路1: 借助暂存空间
//  1. 使用暂存空间，存入指定的slice
func PairSumMethod01(head *ListNode) int {
	// 迭代链表
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	// 求孪生节点最大和
	var maxSum int
	n := len(list)
	for i := 0; i < n/2; i++ {
		curSum := list[i] + list[n-1-i]
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}

func TestPairSum(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want int
	}{
		{"t1", []int{1, 2, 3, 4}, 5},
		{"t2", []int{4, 2, 2, 3}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlink := IntSliceToLinkList(tt.list)
			if got := PairSumMethod02(tlink); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
