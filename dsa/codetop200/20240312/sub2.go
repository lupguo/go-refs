package _0240312

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
// https://leetcode.cn/problems/add-two-numbers/
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	var twoSum, carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int

		// l1 & l2迭代
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// 两数和+进位
		twoSum = v1 + v2 + carry
		remainder := twoSum % 10
		if twoSum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		// 连接新的节点
		p.Next = &ListNode{
			Val:  remainder,
			Next: nil,
		}
		p = p.Next
	}

	// 检测carry进位
	if carry == 1 {
		p.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}
