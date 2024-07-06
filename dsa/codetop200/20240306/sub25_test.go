package _0240306

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// k个一组反转链表:
//
//	head: 1->2->3->4->5 k = 2
//	res: 2->1->4->3->5
//
// https://leetcode.cn/problems/reverse-nodes-in-k-group/
//
//		思路: 迭代+递归
//		  1. p: 反转前k个元素的链表（不足字节返回原始链表）
//		  2. q: 针对k个元素剩余链表进行递归处理，得到反转后的结果
//	   3. 拼接link1, link2
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 异常输入 check
	if head == nil || k == 0 {
		return head
	}

	// base case，链表长度检测，小于k则直接返回
	p := head
	length := 1
	for ; p.Next != nil; p = p.Next {
		length++
	}
	if length < k {
		return head
	}

	// 反转k个链表, 链表头为prev
	var prev *ListNode
	p = head
	length = 0
	for ; p != nil && length < k; length++ {
		tmp := p.Next
		p.Next = prev
		prev = p
		p = tmp
	}

	// 剩余部分递归
	head.Next = reverseKGroup(p, k)

	return prev
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{"t1", []int{1}, 1, []int{1}},
		{"t2", []int{1, 2}, 2, []int{2, 1}},
		{"t3", []int{1, 2, 3}, 2, []int{2, 1, 3}},
		{"t4", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"t5", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := IntSliceToLinkList(tt.input)
			got := reverseKGroup(head, tt.k)
			if !slices.Equal(got.ToIntSlice(), tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}

}
