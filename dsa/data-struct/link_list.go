package data_struct

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) ToIntSlice() []int {
	p := head
	var ret []int
	for p != nil {
		ret = append(ret, p.Val)
		p = p.Next
	}
	return ret
}

// IntLinkListToSlice link list -> []int slice
func IntLinkListToSlice(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

// IntSliceToLinkList []int slice -> link list
func IntSliceToLinkList(list []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, v := range list {
		p.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		p = p.Next
	}

	return head.Next
}
