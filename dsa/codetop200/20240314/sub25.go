package _0240314

import (
	. "dsa/data-struct"
)

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序
func reverseKGroup(head *ListNode, k int) *ListNode {

	// 1. 链表长度小于k，直接返回(base case)
	p := head
	var cnt int
	for p != nil {
		cnt++
		p = p.Next
	}
	if cnt < k {
		return head
	}

	// 2. 反转前k个元素
	//  nil head->head.next
	//  pre cur nxt
	//  pre<-cur nxt
	//  nil<-pre
	//  nil<-pre cur/nxt
	var pre, nxt *ListNode
	cur := head
	for i := 0; i < k; i++ {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	// 3. 拼接后续递归剩余的内容
	// head <- ... pre nxt->
	rkgLeft := reverseKGroup(nxt, k)
	head.Next = rkgLeft

	return pre
}
