package _0240304

import (
	. "dsa/data-struct"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 思路:
//  headAB=headA,headB, headBA=headB,headA
//  p1遍历headAB, p2遍历headBA，检测p1和p2是否相交，返回相交的节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// base case
	if headA == nil || headB == nil {
		return nil
	}

	// 拼接A,B链表
	funcName(headA, headB)

	return nil
}

func funcName(headA *ListNode, headB *ListNode) {
	p := headA
	for p != nil {
		// 如果没有拼接过另外一个链表，则将其拼接起来
		if p.Next == nil {
			p.Next = headB
			break
		}
	}
}
