package _0240127

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
// K个一组反转链表
//
//	基本思路: 递归思路（反转前k个，递归处理子问题，拼接两者）
//	1. 遍历寻找第k个节点（p) ，即(head ->... )k个-> p -> (p.next...)
//	2. 针对head->..p 进行链表反转 p->.. ->head->nil
//	3. 针对p.next进行r = reverseKGroup(p.next, k) 递归迭代，得到r
//	4. head.next = r
//	5. 返回p
//
// 出现问题:
//   - k值的判断 i<k, i=0, k=2, 会遍历两遍
//   - pre,nxt问题， cur是游标，从head到尾nil，最终cur == nxt
//   - base条件，即当反转链表长度小于k时候，直接返回pre值
//   - 迭代反转前k个节点，还是有一些小麻烦（还是应该画清楚图，做好UT）
//   - 如果反转到第k个节点，还是用迭代法好
//   - 没有看清楚题目，如果没有超过k个，则保留原数组（有两种方式：
//   - 一种是遍历一遍链表看下长度是否超过了k
//   - 还一种是遍历i..k，检测i是否已经达到了k
//   - head为0的情况，可以在for循环中cur=nil跳出循环，加i<k的base case中返回，不用首行在包含if (head == nil)...类似判断
func reverseKGroup(head *ListNode, k int) *ListNode {
	// --- 迭代反转前k个节点 ---
	// base case1: 检测是否超过了k个元素，未超过则直接返回原样head
	var i int
	for p := head; i < k; i++ {
		if p == nil { // i未达到k，但链表已经迭代完了，表示链表长度小于k，应该直接返回head
			return head
		}
		p = p.Next
	}

	// base case2: 节点数>=k，则先反转链表前k个节点
	i = 0
	var pre, nxt *ListNode
	for cur := head; i < k && cur != nil; i++ {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 恰好k个元素，直接返回pre反转链表
	if nxt == nil {
		return pre
	}

	// --- 迭代反转前k个节点 end ---

	// 现在cur指向第k个节点，nxt指向k+1个节点
	// 递归处理子链表
	revKG := reverseKGroup(nxt, k)
	head.Next = revKG

	return pre
}

// 寻找到[0,...k) 节点 k = p, 反转[head,p)
func reverseKGroupV2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// 找到p
	p := head
	for i := 0; i < k; i++ {
		if p == nil { // i没超过k，即表示链表head没有超过k，所以直接返回head
			return head
		}
		p = p.Next
	}
	// 反转[head, p)
	ret := reverseHeadToNode(head, p)
	// 将head.Next(反转最后一个元素的下个节点）设置为递归KGroup(p.next)后的链表
	head.Next = reverseKGroupV2(p, k)

	return ret
}

// 迭代反转链表[head, p), 从头到p位置节点
// p == nil 表示要反转整个head
func reverseHeadToNode(head, p *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre, nxt *ListNode
	for cur := head; cur != p; {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	// head拼接
	head.Next = nxt
	return pre
}

func TestReverseHeadToNode(t *testing.T) {
	// 创建链表节点
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	// 构建链表：1 -> 2 -> 3
	node1.Next = node2
	node2.Next = node3

	tests := []struct {
		name   string
		head   *ListNode
		p      *ListNode
		output []int
	}{
		{"t1", node1, node2, []int{2, 1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseHeadToNode(tt.head, tt.p)

			// 验证反转后的链表节点值是否与预期一致
			for i, val := range tt.output {
				if got.Val != val {
					t.Errorf("got value %v at index %v, but want %v", got.Val, i, val)
				}
				got = got.Next
			}
		})
	}
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
		{"t3", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"t4", []int{1, 2, 3}, 2, []int{2, 1, 3}},
		{"t5", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlist := IntSliceToLinkList(tt.input)
			tgot := reverseKGroupV2(tlist, tt.k).ToIntSlice()
			if !slices.Equal(tgot, tt.want) {
				t.Errorf("got %v, but want %v", tgot, tt.want)
			}
		})
	}
}
