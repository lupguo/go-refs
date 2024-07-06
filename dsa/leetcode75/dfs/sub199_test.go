package dfs

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{root: nil}, nil},
		{"t2", args{root: &TreeNode{
			Val: 100,
		}}, []int{100}},
		{"t3", args{root: &TreeNode{
			Val: 100,
			Right: &TreeNode{
				Val: 200,
			},
		}}, []int{100, 200}},
		{"t4", args{root: &TreeNode{
			Val: 100,
			Left: &TreeNode{
				Val: 200,
			},
		}}, []int{100, 200}},
		{"t5", args{root: &TreeNode{
			Val: 100,
			Left: &TreeNode{
				Val: 200,
				Right: &TreeNode{
					Val: 300,
				},
			},
		}}, []int{100, 200, 300}},
		{"t6", args{root: &TreeNode{
			Val: 100,
			Right: &TreeNode{
				Val: 200,
				Right: &TreeNode{
					Val: 300,
				},
			},
		}}, []int{100, 200, 300}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
