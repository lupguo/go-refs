package list

import (
	"testing"
)

func TestCreateAndPrintList(t *testing.T) {
	head := CreateAndPrintList()
	// t.Log(head.Next.ToSlice())
	//
	// rev := reverseList2(head.Next)
	// t.Log(rev.ToSlice())

	// 递归反向打印
	t.Log(head.Next.Reverse())
}
