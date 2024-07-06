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

// 给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。
//
// 长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。
//
// 对于 n = 1、2、3、4 和 5 的情况，中间节点的下标分别是 0、1、1、2 和 2 。

// 示例:
// 输入：head = [1,3,4,7,1,2,6]
// 输出：[1,3,4,1,2,6]
// 解释：
// 上图表示给出的链表。节点的下标分别标注在每个节点的下方。
// 由于 n = 7 ，值为 7 的节点 3 是中间节点，用红色标注。
// 返回结果为移除节点后的新链表。
func deleteMiddle(head *ListNode) *ListNode {

	return nil
}

// 思路2：快慢指针处理
//  1. 申请两个指针fast,slow，快指针fast尝试走两步，slow走一步:
//  2. 异常排除: 如果fast ==nil || fast.next == nil return nil
//  3. 迭代链表:
//     * 如果 fast.next != nil && fast.next.next != nil { fast = fast.next.next; slow = slow.next }
//  4. 迭代完后slow的下个节点为待删除的前驱节点 slow.next = slow.next.next
func DeleteMiddleMethod02(head *ListNode) *ListNode {
	fast, slow := head, head
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 有2个或2个以上节点的迭代链表
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		if fast.Next != nil { // 这里需要小心检查下，快指针是否走到头了，如果走到头了，慢指针则不移动
			slow = slow.Next
		}
	}

	// 删除slow指针的后继节点
	slow.Next = slow.Next.Next

	return head
}

// 思路:
//
//	统计链表长度，找中间节点前驱节点，删除处理
//	1. 从前往后迭代链表 i=0,1,2...n-1 -> 你无法指定链表长度，先统计一次
//	2. 检测当前节点是否为中间节点(i==n/2)，找到该节点cur，尝试删除该cur节点 。因为删除cur节点，依赖与cur的前驱节点，所以条件判断更改成了
//	   当 i+1 = n/2时候，删除cur.next，就变成了cur.next = cur.next.next
//	3. 无需检测cur.next是否为nil，因为我们是从i节点开始统计，i+1节点一定会存在，在i+1=n/2时候，cur.next一定不会为空，所以cur.next.next也不会报错
func DeleteMiddle(head *ListNode) *ListNode {
	// 统计链表长度
	var llen int
	for p := head; p != nil; p = p.Next {
		llen++
	}

	// 只有一个节点，删除后就为空链表了
	if llen == 1 {
		return nil
	}

	// 有超过一个节点，重新迭代链表，在中间节点前一个节点停止
	var i int
	for cur := head; cur != nil; i++ {
		if i+1 == llen/2 { // 找到中间节点前驱节点
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}

	return head
}

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"t1", []int{1, 2, 3, 4}, []int{1, 2, 4}},
		{"t2", []int{2, 1}, []int{2}},
		{"t3", []int{1}, nil},
		{"t4", []int{1, 2, 3}, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlist := IntSliceToLinkList(tt.list)
			dtlink := DeleteMiddleMethod02(tlist)
			gotList := IntLinkListToSlice(dtlink)

			if !slices.Equal(gotList, tt.want) {
				t.Errorf("got %v, but want %v", gotList, tt.want)
			}
		})
	}
}
