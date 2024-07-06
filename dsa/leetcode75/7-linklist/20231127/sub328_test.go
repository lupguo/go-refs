package _0231127

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
//
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
func oddEvenList(head *ListNode) *ListNode {

	return nil
}

// 思路: 组建两条链，让基数链的链尾连上偶数连的链头
//  1. 初始化op和ep两个指针迭代，组成两条新的链 opHead和epHead，让迭代到最后
//  2. ep迭代偶数链(ep.next = ep.next.next, ep = ep.next), op迭代基数链（同样)
//  3. 最后串联奇数和偶数链: op.next = epHead, 同时注意ep.next = nil
//  4. 返回opHead
func OddEvenList(head *ListNode) *ListNode {
	// 小于1个节点情况
	if head == nil || head.Next == nil {
		return head
	}

	// 大于1个节点情况，初始奇数、偶数链
	opHead, epHead := head, head.Next
	op, ep := head, head.Next

	for op.Next != nil && op.Next.Next != nil {
		// 先组建奇数链
		op.Next = op.Next.Next
		op = op.Next

		// 再组建偶数链
		if ep.Next.Next != nil {
			ep.Next = ep.Next.Next
			ep = ep.Next
		}
	}

	// 串联奇、偶链
	op.Next = epHead
	ep.Next = nil

	return opHead
}

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"t1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{"t2", []int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlinkList := IntSliceToLinkList(tt.list)
			slist := OddEvenList(tlinkList)
			got := IntLinkListToSlice(slist)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
