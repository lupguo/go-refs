package _0231127

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	return nil
}

// 思路: 循环
// 前缀指针记录
func ReverseListLoop(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	for head.Next != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}

	head.Next = prev
	return head
}

// 思路: 递归
//  1. 假定只有一个元素或者空元素情况, base条件
//  2. 假定 ReverseListRec(head.Next) 已反转完成了
func ReverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := ReverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{"t1", []int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlink := IntSliceToLinkList(tt.head)
			rlink := ReverseListLoop(tlink)
			got := IntLinkListToSlice(rlink)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
