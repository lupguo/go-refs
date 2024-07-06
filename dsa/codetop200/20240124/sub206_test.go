package _0240124

import (
	"slices"
	"testing"
	"unsafe"

	. "dsa/data-struct"
)

// 反转链表
//
//	思路: 递归思路，基准条件判断、后续遍历位置、利用递归定义
func reverseList(head *ListNode) *ListNode {
	// base情况
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	// 递归处理
	revHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return revHead
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"t1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"t2", []int{1}, []int{1}},
		{"t3", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := IntSliceToLinkList(tt.input)
			got := IntLinkListToSlice(reverseList(head))

			if !slices.Equal(got, tt.output) {
				t.Errorf("got %v, but want %v", got, tt.output)
			}
		})
	}
}

func TestNil(t *testing.T) {
	var slice1 []int
	slice2 := []int{}
	var slice3 []int = nil

	t.Logf("var []int size:%v", unsafe.Sizeof(slice1))
	t.Logf("[]int{} size:%v", unsafe.Sizeof([]int{}))
	t.Logf("[]int{} size:%v", unsafe.Sizeof(slice3))

	t.Log(slice1) // []
	t.Log(slice2) // []
	t.Log(slice3) // []

	t.Log(len(slice1)) // 0
	t.Log(len(slice2)) // 0
	t.Log(len(slice3)) // 0

	// 添加元素到切片
	slice1 = append(slice1, 1, 2, 3)
	slice2 = append(slice2, 1, 2, 3)
	slice3 = append(slice3, 1, 2, 3)

	t.Log(cap(slice1)) // 0
	t.Log(cap(slice2)) // 0
	t.Log(cap(slice3)) // 0

	t.Log(slice1) // [1 2 3]
	t.Log(slice2) // []
	t.Log(slice3) // []
}
