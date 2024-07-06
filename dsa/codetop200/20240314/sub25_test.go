package _0240314

import (
	"slices"
	"testing"

	. "dsa/data-struct"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1}, 1}, []int{1}},
		{"t2", args{[]int{1, 2}, 1}, []int{1, 2}},
		{"t3", args{[]int{1, 2}, 2}, []int{2, 1}},
		{"t4", args{[]int{1, 2}, 3}, []int{1, 2}},
		{"t5", args{[]int{1, 2, 3}, 2}, []int{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := IntSliceToLinkList(tt.args.nums)
			gotList := reverseKGroup(head, tt.args.k)
			got := gotList.ToIntSlice()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
