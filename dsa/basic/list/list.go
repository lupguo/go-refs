package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// ToSlice ListNode转Slice
func (l *ListNode) ToSlice() []int {
	var list []int
	for ; l != nil; l = l.Next {
		list = append(list, l.Val)
	}
	return list
}

func (l *ListNode) Reverse() []int {
	// 终止条件
	if l == nil {
		return nil
	}

	// 递归调用
	s := l.Next.Reverse()

	// 逻辑处理
	s = append(s, l.Val)

	return s
}

func CreateAndPrintList() *ListNode {
	// 创建链表
	head := &ListNode{}
	cur := head // 用一个指针来代表链表当前的头节点
	for _, v := range []int{1, 2, 3} {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	// 遍历链表
	cur = head.Next
	for cur != nil {
		cur = cur.Next
	}

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用三个指针 pre、cur、next，分别指向当前节点的前一个节点、当前节点、当前节点的下一个节点。
// 我们从头节点开始遍历链表，逐个反转链表中的节点，反转操作的过程中，需要将当前节点的 next 指针指向前一个节点，然后依次移动三个指针。
func reverseList1(head *ListNode) *ListNode {
	var rList *ListNode
	if head.Next != nil {
		rList = reverseList1(head.Next)
		// head -> head.Next -> head.Next.Next ...
		// head <- head.Next ... <- rList
		head.Next = head
		return rList
	}

	return head
}

func reverseList2(head *ListNode) *ListNode {
	// 压入slice
	var nodes []*ListNode
	for {
		if head.Next != nil {
			nodes = append(nodes, head)
		} else {
			nodes = append(nodes, head)
			break
		}
		head = head.Next
	}

	// 从slice最后一个元素读取，如果i > 1, 则依次指向前一个元素
	rev := nodes[len(nodes)-1]
	for i := len(nodes) - 1; i >= 0; i-- {
		if i > 0 {
			nodes[i].Next = nodes[i-1]
		} else {
			nodes[i].Next = nil
		}
	}

	return rev
}

// 两个链表
// 1->2->3->nil
// nil(rev)
//
// 2->3->nil
// 1-> nil
// PASS:
//
// 3->nil
// 2->1->nil
func reverseList3(head *ListNode) *ListNode {
	var newLink *ListNode
	for head != nil {
		// 暂存head.Next防止后继丢失
		temp := head.Next

		// head.Next现在指向新的链表
		head.Next = newLink

		// newLink也要更新成最新的head
		newLink = head

		// 原链表现在可以向后迭代了
		head = temp
	}

	return newLink
}

func reverseList(head *ListNode) *ListNode {
	// 终止条件
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转head.Next
	rev := reverseList(head.Next)

	// 逻辑处理 head.Next的Next指针，要指向回head
	head.Next.Next = head

	// head.Next指针要置位空，不然会有环存在
	head.Next = nil

	return rev
}
