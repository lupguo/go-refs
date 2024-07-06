package _0240312

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{l1: IntSliceToLinkList([]int{1, 2}), l2: IntSliceToLinkList([]int{1, 2})}, []int{2, 4}},
		{"t2", args{l1: IntSliceToLinkList([]int{1}), l2: IntSliceToLinkList([]int{1, 2})}, []int{2, 2}},
		{"t3", args{l1: IntSliceToLinkList([]int{9}), l2: IntSliceToLinkList([]int{1, 2})}, []int{0, 3}},
		{"t4", args{l1: IntSliceToLinkList([]int{9}), l2: IntSliceToLinkList([]int{1, 9})}, []int{0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList := addTwoNumbers(tt.args.l1, tt.args.l2)
			got := gotList.ToIntSlice()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
